package utils

import (
	"errors"
	"log/slog"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func IdsToStringArray(ids []primitive.ObjectID) []string {
	result := make([]string, len(ids))

	for i, id := range ids {
		result[i] = id.Hex()
	}

	return result
}

func StringsToIdsArray(ids []string) ([]primitive.ObjectID, error) {
	result := make([]primitive.ObjectID, len(ids))

	for i, id := range ids {
		newId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			slog.Error("[StringsToIdsArray] error: " + err.Error())
			return nil, errors.New("[StringsToIdsArray] error: " + err.Error())
		}
		result[i] = newId
	}

	return result, nil
}
