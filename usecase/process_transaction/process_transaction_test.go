package process_transaction

import (
	"testing"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/jvcalassio/fc-payment-gateway/domain/repository/mock"
	"github.com/jvcalassio/fc-payment-gateway/adapter/broker/mock"
	"github.com/jvcalassio/fc-payment-gateway/domain/entity"
)

func TestProcessTransaction_InvalidCC(t *testing.T) {
	input := TransactionDtoInput{
		ID: "1",
		AccountID: "1",
		CreditCardNumber: "999999999999999",
		CreditCardName: "Joao Silva",
		CreditCardExpMonth: 12,
		CreditCardExpYear: 2027,
		CreditCardCVV: 999,
		Amount: 200,
	}

	expectedOutput := TransactionDtoOutput{
		ID: "1",
		Status: entity.REJECTED,
		ErrorMessage: "invalid credit card number",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().
		Insert(input.ID, input.AccountID, input.Amount, expectedOutput.Status, expectedOutput.ErrorMessage).
		Return(nil)

	producerMock := mock_broker.NewMockProducerInterface(ctrl)
	producerMock.EXPECT().
		Publish(expectedOutput, []byte(input.ID), "transactions_result")

	usecase := NewProcessTransaction(repositoryMock, producerMock, "transactions_result")

	output, err := usecase.Execute(input)
	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}

func TestProcessTransaction_Rejected(t *testing.T) {
	input := TransactionDtoInput{
		ID: "1",
		AccountID: "1",
		CreditCardNumber: "5438330212460978",
		CreditCardName: "Joao Silva",
		CreditCardExpMonth: 12,
		CreditCardExpYear: 2027,
		CreditCardCVV: 999,
		Amount: 1200,
	}

	expectedOutput := TransactionDtoOutput{
		ID: "1",
		Status: entity.REJECTED,
		ErrorMessage: "no limit for this transaction",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().
		Insert(input.ID, input.AccountID, input.Amount, expectedOutput.Status, expectedOutput.ErrorMessage).
		Return(nil)

	producerMock := mock_broker.NewMockProducerInterface(ctrl)
	producerMock.EXPECT().
		Publish(expectedOutput, []byte(input.ID), "transactions_result")

	usecase := NewProcessTransaction(repositoryMock, producerMock, "transactions_result")

	output, err := usecase.Execute(input)
	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}

func TestProcessTransaction_Approved(t *testing.T) {
	input := TransactionDtoInput{
		ID: "1",
		AccountID: "1",
		CreditCardNumber: "5438330212460978",
		CreditCardName: "Joao Silva",
		CreditCardExpMonth: 12,
		CreditCardExpYear: 2027,
		CreditCardCVV: 999,
		Amount: 700,
	}

	expectedOutput := TransactionDtoOutput{
		ID: "1",
		Status: entity.APPROVED,
		ErrorMessage: "",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().
		Insert(input.ID, input.AccountID, input.Amount, expectedOutput.Status, expectedOutput.ErrorMessage).
		Return(nil)

	producerMock := mock_broker.NewMockProducerInterface(ctrl)
	producerMock.EXPECT().
		Publish(expectedOutput, []byte(input.ID), "transactions_result")

	usecase := NewProcessTransaction(repositoryMock, producerMock, "transactions_result")

	output, err := usecase.Execute(input)
	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}