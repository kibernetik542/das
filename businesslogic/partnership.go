package businesslogic

import (
	"errors"
	"time"
)

const (
	// PartnershipRequestReceived labels the request as "received", once it's viewed by recipient
	PartnershipRequestReceived = 1
	// PartnershipRequestSent labels the request at "sent", once it's sent out by sender
	PartnershipRequestSent = 2
)

const (
	// PartnershipRoleLead is the reference value for the Lead role
	PartnershipRoleLead = 2
	// PartnershipRoleFollow is the reference value for the Follow role
	PartnershipRoleFollow = 1
)

// PartnershipRole defines roles within a Partnership: lead and follow
type PartnershipRole struct {
	ID              int
	Name            string
	DateTimeCreated time.Time
	DateTimeUpdated time.Time
}

// IPartnershipRoleRepository specifies the interface that a PartnershipRoleRepository should implement
type IPartnershipRoleRepository interface {
	GetAllPartnershipRoles() ([]PartnershipRole, error)
}

type PartnershipStatus struct {
	ID              int
	Name            string
	DateTimeCreated time.Time
	DateTimeUpdated time.Time
}

type IPartnershipStatusRepository interface {
	GetAllPartnershipStatus() ([]PartnershipStatus, error)
}

// Partnership defines the combination of a lead and a follow. A partnership is uniquely identified
// if the lead and follow are confirmed.
type Partnership struct {
	ID                   int
	Lead                 Account
	Follow               Account
	SameSex              bool
	FavoriteByLead       bool
	FavoriteByFollow     bool
	CompetitionsAttended int
	EventsAttended       int
	DateTimeCreated      time.Time
	DateTimeUpdated      time.Time
}

// IPartnershipRepository defines the interface that a partnership repository should implement
type IPartnershipRepository interface {
	CreatePartnership(partnership *Partnership) error
	SearchPartnership(criteria SearchPartnershipCriteria) ([]Partnership, error)
	UpdatePartnership(partnership Partnership) error
	DeletePartnership(partnership Partnership) error
}

// SearchPartnershipCriteria provides the parameters that an IPartnershipRepository can use to search by
type SearchPartnershipCriteria struct {
	PartnershipID int `schema:"id"`
	LeadID        int `schema:"lead"`
	FollowID      int `schema:"follow"`
	AccountID     int
}

// GetAllPartnerships returns all the partnerships that caller account is in, including as a lead and as a follow
func (account Account) GetAllPartnerships(repo IPartnershipRepository) ([]Partnership, error) {
	asLeads, err := repo.SearchPartnership(SearchPartnershipCriteria{
		LeadID: account.ID,
	})
	if err != nil {
		return nil, err
	}

	asFollows, err := repo.SearchPartnership(SearchPartnershipCriteria{
		FollowID: account.ID,
	})
	if err != nil {
		return nil, err
	}

	allPartnerships := make([]Partnership, 0)
	for _, each := range asLeads {
		allPartnerships = append(allPartnerships, each)
	}
	for _, each := range asFollows {
		allPartnerships = append(allPartnerships, each)
	}
	return allPartnerships, err
}

// GetPartnershipByID retrieves the Partnership in the provided repository by the specified ID
func GetPartnershipByID(id int, repo IPartnershipRepository) (Partnership, error) {
	searchResults, err := repo.SearchPartnership(SearchPartnershipCriteria{PartnershipID: id})
	if err != nil || searchResults == nil || len(searchResults) != 1 {
		return Partnership{}, err
	}
	return searchResults[0], err
}

// MustGetPartnershipByID uses an known ID and a concrete PartnershipRepository to find the
// partnership by the ID provided. If such partnership is not found, system will panic.
func MustGetPartnershipByID(id int, repo IPartnershipRepository) Partnership {
	searchResults, err := repo.SearchPartnership(SearchPartnershipCriteria{PartnershipID: id})
	if err != nil {
		panic(err.Error())
	}
	if len(searchResults) != 1 {
		panic("cannot find partnership with this ID")
	}
	return searchResults[0]
}

// HasAthlete checks if the provided Athlete ID is in this partnership
func (partnership Partnership) HasAthlete(athleteID int) bool {
	return partnership.Lead.ID == athleteID || partnership.Follow.ID == athleteID
}

const (
	// PartnershipRequestStatusAccepted is the status of a request when it is accepted by the recipient
	PartnershipRequestStatusAccepted = 1
	// PartnershipRequestStatusPending is the status of a request when it is sent out, but not responded by recipient
	PartnershipRequestStatusPending = 2
	// PartnershipRequestStatusDeclined is the status of a request when it is declined by the recipient
	PartnershipRequestStatusDeclined = 3
)

type PartnershipRequestStatus struct {
	ID              int
	Code            string
	Description     string
	DateTimeCreated time.Time
	DateTimeUpdated time.Time
}

type IPartnershipRequestStatusRepository interface {
	GetPartnershipRequestStatus() ([]PartnershipRequestStatus, error)
}

// PartnershipRequest is a request of partnership that can only be sent between Athletes who are not in the requested Partnership
type PartnershipRequest struct {
	PartnershipRequestID int
	SenderID             int
	RecipientID          int
	SenderAccount        *Account
	RecipientAccount     *Account
	SenderRole           int
	RecipientRole        int
	Message              string
	Status               int
	CreateUserID         int
	DateTimeCreated      time.Time
	UpdateUserID         int
	DateTimeUpdated      time.Time
}

// PartnershipRequestResponse is the response that recipient can make towards a request
type PartnershipRequestResponse struct {
	RequestID       int
	RecipientID     int
	Response        int
	DateTimeCreated time.Time
}

// SearchPartnershipRequestCriteria defines the parameters that can be used to search particular partnership requests
type SearchPartnershipRequestCriteria struct {
	RequestID       int `schema:"id"`
	Type            int `schema:"typeId"`
	Sender          int `schema:"sender"`
	Recipient       int `schema:"recipient"`
	RequestStatusID int `schema:"statusId"`
}

// IPartnershipRequestRepository specifies the functions that need to be implemented to allow CRUD operations on
// Partnership Request
type IPartnershipRequestRepository interface {
	CreatePartnershipRequest(request *PartnershipRequest) error
	SearchPartnershipRequest(criteria SearchPartnershipRequestCriteria) ([]PartnershipRequest, error)
	DeletePartnershipRequest(request PartnershipRequest) error
	UpdatePartnershipRequest(request PartnershipRequest) error
	GetReceivedRequests(recipientID int) ([]PartnershipRequest, error)
	GetSentRequests(senderID int) ([]PartnershipRequest, error)
}

type PartnershipRequestService struct {
	accountRepo     IAccountRepository
	partnershipRepo IPartnershipRepository
	blacklistRepo   IPartnershipRequestBlacklistRepository
}

func NewPartnershipRequestService(accountRepo IAccountRepository, partnershipRepo IPartnershipRepository, blacklistRepo IPartnershipRequestBlacklistRepository) *PartnershipRequestService {
	service := new(PartnershipRequestService)
	service.accountRepo = accountRepo
	service.partnershipRepo = partnershipRepo
	service.blacklistRepo = blacklistRepo
	return service
}

func (service PartnershipRequestService) CreatePartnershipRequest(currentUser Account, request PartnershipRequest) error {
	if currentUser.ID != request.SenderID {
		return errors.New("not authorized to send this partnership request")
	}
	return errors.New("not implemented")
}

func (request PartnershipRequest) validateRoles() error {
	if request.SenderRole != PartnershipRoleLead && request.SenderRole != PartnershipRoleFollow {
		return errors.New("sender's role is not specified")
	}
	if request.RecipientRole != PartnershipRoleLead && request.RecipientRole != PartnershipRoleFollow {
		return errors.New("recipient's role is not specified")
	}
	if request.RecipientRole == request.SenderRole {
		return errors.New("sender and recipient have identical roles")
	}
	if request.SenderID == request.RecipientID {
		return errors.New("cannot send partnership request to yourself")
	}
	return nil
}

func (request *PartnershipRequest) hasValidSenderAndRecipient(accountRepo IAccountRepository) error {
	senderAccounts, seErr := accountRepo.SearchAccount(SearchAccountCriteria{ID: request.SenderID})
	recipientAccounts, recErr := accountRepo.SearchAccount(SearchAccountCriteria{ID: request.RecipientID})
	if seErr != nil {
		return seErr
	}
	if recErr != nil {
		return recErr
	}
	if len(senderAccounts) != 1 {
		return errors.New("sender account cannot be found")
	}
	if len(recipientAccounts) != 1 {
		return errors.New("recipient account cannot be found")
	}
	sender := senderAccounts[0]
	recipient := recipientAccounts[0]

	if !sender.HasRole(AccountTypeAthlete) {
		return errors.New("sender is not an athlete")
	}
	if !recipient.HasRole(AccountTypeAthlete) {
		return errors.New("recipient is not an athlete")
	}

	request.SenderAccount = &senderAccounts[0]
	request.RecipientAccount = &recipientAccounts[0]
	return nil
}

func (request PartnershipRequest) senderBlockedByRecipient(blacklistRepo IPartnershipRequestBlacklistRepository) bool {
	recipientBlacklist, _ := blacklistRepo.SearchPartnershipRequestBlacklist(SearchPartnershipRequestBlacklistCriteria{ReporterID: request.RecipientID})
	for _, each := range recipientBlacklist {
		if each.BlockedUser.ID == request.SenderID {
			return true
		}
	}
	return false
}

// hasExistingPartnership checks if there is already a partnership between the two dancers
func (request PartnershipRequest) hasExistingPartnership(accountRepo IAccountRepository, partnershipRepo IPartnershipRepository) bool {
	// configure search partnershipCriteria

	var senderAccount = *request.SenderAccount
	var recipientAccount = *request.RecipientAccount
	if request.SenderAccount == nil {
		senderAccount = GetAccountByID(request.SenderID, accountRepo)
	}
	if request.RecipientAccount == nil {
		recipientAccount = GetAccountByID(request.RecipientID, accountRepo)
	}

	partnershipCriteria := new(SearchPartnershipCriteria)
	if request.SenderRole == PartnershipRoleLead {
		partnershipCriteria.LeadID = senderAccount.ID
		partnershipCriteria.FollowID = recipientAccount.ID
	} else {
		partnershipCriteria.FollowID = senderAccount.ID
		partnershipCriteria.LeadID = recipientAccount.ID
	}

	// check if sender is already in a partnership with recipient
	partnerships, _ := partnershipRepo.SearchPartnership(*partnershipCriteria)
	if len(partnerships) != 0 {
		return true
	}
	return false
}

// hasPendingRequest checks if there is a request between these two dancers that still waits for response
func (request PartnershipRequest) hasPendingRequest(requestRepo IPartnershipRequestRepository) bool {
	// check if there is pending message between sender and recipient
	requests, _ := requestRepo.SearchPartnershipRequest(SearchPartnershipRequestCriteria{
		Recipient:       request.RecipientID,
		Sender:          request.SenderID,
		RequestStatusID: PartnershipRequestStatusPending,
	})
	if len(requests) == 1 {
		return true
	}
	return false
}

// CreatePartnershipRequest will create the partnership request with validation. Validation includes
// 1. Role validation: must be opposite role
// 2. Blacklist check: sender must not be blacklisted by recipient
// 3. Existing partnership check: sender and recipient must not be in a partnership with specified role
// 4. There is no pending request for the same role (this is applied to request from either party)
// Note: if sender and recipient are in a partnership of opposite role, then it's considered as a different partnership.
// If the request is valid, then request will be created.
func CreatePartnershipRequest(request PartnershipRequest, partnershipRepo IPartnershipRepository,
	requestRepo IPartnershipRequestRepository, accountRepo IAccountRepository,
	blacklistRepo IPartnershipRequestBlacklistRepository) error {

	// validate Roles the request first
	if roleErr := request.validateRoles(); roleErr != nil {
		return roleErr
	}

	// check if accounts exist
	if accountErr := request.hasValidSenderAndRecipient(accountRepo); accountErr != nil {
		return accountErr
	}

	// check if sender is blacklisted by recipient
	if request.senderBlockedByRecipient(blacklistRepo) {
		return errors.New("cannot send partnership request to this user")
	}

	if request.hasExistingPartnership(accountRepo, partnershipRepo) {
		return errors.New("you are already in a partnership with specified role")
	}

	if request.hasPendingRequest(requestRepo) {
		return errors.New("a pending request must be responded first")
	}

	return requestRepo.CreatePartnershipRequest(&request)
}

func validatePartnershipRequestResponse(response PartnershipRequestResponse, repo IPartnershipRequestRepository) error {
	if response.RecipientID == 0 {
		return errors.New("recipient must be specified")
	}

	if response.RequestID == 0 {
		return errors.New("request must be specified")
	}

	// check if request is valid
	if requests, searchErr := repo.SearchPartnershipRequest(SearchPartnershipRequestCriteria{
		RequestID: response.RequestID,
		Recipient: response.RecipientID,
	}); searchErr != nil {
		return searchErr
	} else if len(requests) != 1 {
		return errors.New("cannot find request for this recipient")
	} else if requests[0].Status == PartnershipRequestStatusAccepted || requests[0].Status == PartnershipRequestStatusDeclined {
		return errors.New("this request is already responded")
	}

	return nil
}

func RespondPartnershipRequest(response PartnershipRequestResponse,
	requestRepo IPartnershipRequestRepository,
	accountRepo IAccountRepository,
	partnershipRepo IPartnershipRepository) error {

	if validErr := validatePartnershipRequestResponse(response, requestRepo); validErr != nil {
		return validErr
	}

	// respond partnership
	if response.Response == PartnershipRequestStatusAccepted || response.Response == PartnershipRequestStatusDeclined {
		requests, err := requestRepo.SearchPartnershipRequest(SearchPartnershipRequestCriteria{
			RequestID: response.RequestID,
			Recipient: response.RecipientID,
		})
		if err != nil {
			return err
		}
		requests[0].DateTimeUpdated = time.Now()
		requests[0].Status = response.Response
		if respErr := requestRepo.UpdatePartnershipRequest(requests[0]); respErr != nil {
			return respErr
		}

		// optional: create partnership if accepted
		if response.Response == PartnershipRequestStatusAccepted {
			partnership := Partnership{}
			requests, _ := requestRepo.SearchPartnershipRequest(SearchPartnershipRequestCriteria{
				RequestID: response.RequestID,
				Recipient: response.RecipientID,
			})
			request := requests[0]

			if request.RecipientRole == PartnershipRoleLead {
				partnership.Lead.ID = request.RecipientID
				partnership.Follow.ID = request.SenderID
			} else {
				partnership.Lead.ID = request.SenderID
				partnership.Follow.ID = request.RecipientID
			}

			leadAccount := GetAccountByID(partnership.Lead.ID, accountRepo)
			followAccount := GetAccountByID(partnership.Follow.ID, accountRepo)
			if leadAccount.UserGenderID == followAccount.UserGenderID {
				partnership.SameSex = true
			} else {
				partnership.SameSex = false
			}

			partnership.DateTimeCreated = time.Now()
			return partnershipRepo.CreatePartnership(&partnership)
		}
	}
	return nil
}
