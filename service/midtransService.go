package service

import (
	"booking-konstruksi/repository"
	"booking-konstruksi/request"
	"booking-konstruksi/response"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"os"
)

type MidtransService interface {
	Create(request request.MidtransRequest, id string) (response.MidtransResponse, error)
}

type midtransService struct {
	repoPembayaran repository.Pembayaran
}

func NewMidtransService(repoPembayaran repository.Pembayaran) *midtransService {
	return &midtransService{
		repoPembayaran: repoPembayaran,
	}
}

func (s *midtransService) Create(request request.MidtransRequest, id string) (response.MidtransResponse, error) {
	// request midtrans
	var snapClient = snap.Client{}
	snapClient.New(os.Getenv("MIDTRANS_SERVER_KEY"), midtrans.Sandbox)

	// pembayaran
	pembayaran, err := s.repoPembayaran.GetData(id)
	amount, _ := request.Amount.Int64()

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  pembayaran.Kode,
			GrossAmt: amount,
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: pembayaran.Client.Name,
			Email: pembayaran.Client.Email,
			Phone: pembayaran.Client.NoWA,
		},
		EnabledPayments: []snap.SnapPaymentType{
			snap.PaymentTypeBankTransfer,
		},
		Items: &[]midtrans.ItemDetails{
			{
				ID:    "Property-" + request.ItemID,
				Qty:   1,
				Price: amount,
				Name:  request.ItemName,
			},
		},
	}

	response_, errSnap := snapClient.CreateTransaction(req)
	if errSnap != nil {
		panic(errSnap.Error())
	}

	midtransReponse := response.MidtransResponse{
		Token:       response_.Token,
		RedirectUrl: response_.RedirectURL,
	}

	pembayaran.Token = response_.Token
	_, err = s.repoPembayaran.Update(pembayaran)

	return midtransReponse, err
}
