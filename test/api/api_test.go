package api

import "testing"

func TestAll(t *testing.T) {
	t.Run("User Test", func(t *testing.T) {
		TestPostUserWithValidPayload(t)
		TestPostUserWithInvalidPayload(t)
		TestVerifyUserWithValidPayload(t)
		TestVerifyUserWithInvalidPayload(t)
		TestGetUserByIdWithToken(t)
		TestGetUserByIdWithoutToken(t)
		TestUpdateUserByIdWithValidPayload(t)
		TestUpdateUserByIdWithInvalidPayload(t)
	})

	t.Run("Product Test", func(t *testing.T) {
		TestPostProductWithoutImage(t)
		TestPostProductWithImage(t)
		TestPostProductWithInvalidPayload(t)
		TestGetProductsByUserWithToken(t)
		TestGetProductsByUserWithoutToken(t)
		TestGetProductByIdWithValidId(t)
		TestGetProductByIdWithInvalidId(t)
		TestUpdateProductByIdWithoutImage(t)
		TestUpdateProductByIdWithImage(t)
		TestUpdateProductByIdWithInvalidPayload(t)
		TestDeleteProductById(t)
	})
}
