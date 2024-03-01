package dynamoDb

import (
	"context"
	"fmt"

	"github.com/Kenmuraki5/kro-backend.git/domain/entity"
	"github.com/Kenmuraki5/kro-backend.git/domain/restmodel"
	"github.com/google/uuid"

	// "github.com/Kenmuraki5/kro-backend.git/domain/repository"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type DynamoDBOrderRepository struct {
	Client *dynamodb.Client
}

func NewDynamoDBOrderRepository(client *dynamodb.Client) *DynamoDBOrderRepository {
	return &DynamoDBOrderRepository{Client: client}
}

func (repo *DynamoDBOrderRepository) GetAllOrders() ([]*entity.Order, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String("Orders"),
	}
	result, err := repo.Client.Scan(context.TODO(), input)
	if err != nil {
		return nil, fmt.Errorf("failed to scan DynamoDB table: %v", err)
	}
	var Orders []*entity.Order
	for _, item := range result.Items {
		fmt.Println(item)
		var Order entity.Order
		err := attributevalue.UnmarshalMap(item, &Order)
		if err != nil {
			return nil, err
		}
		Orders = append(Orders, &Order)
	}
	fmt.Println(Orders)
	return Orders, nil
}

func (repo *DynamoDBOrderRepository) AddOrders(orders []restmodel.Order) ([]*restmodel.Order, error) {
	var addedOrders []*restmodel.Order
	var writeRequests []types.WriteRequest
	orderID := uuid.NewString()
	for _, order := range orders {

		entityOrder := entity.Order{
			OrderId:         orderID,
			ProductId:       order.ProductId,
			Quantity:        order.Quantity,
			CustomerId:      order.CustomerId,
			OrderDate:       order.OrderDate,
			Status:          "Pending",
			Subtotal:        order.Subtotal,
			ShippingAddress: order.ShippingAddress,
			Type:            order.Type,
		}

		item, err := attributevalue.MarshalMap(entityOrder)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal Order to DynamoDB attribute map: %v", err)
		}

		writeRequest := types.WriteRequest{
			PutRequest: &types.PutRequest{
				Item: item,
			},
		}

		writeRequests = append(writeRequests, writeRequest)
		addedOrders = append(addedOrders, &order)
	}

	_, err := repo.Client.BatchWriteItem(context.TODO(), &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]types.WriteRequest{
			"Orders": writeRequests,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to add Orders to DynamoDB: %v", err)
	}

	return addedOrders, nil
}

// update
func (repo *DynamoDBOrderRepository) UpdateOrder(updatedOrder entity.Order) (*entity.Order, error) {
	if updatedOrder.OrderId == "" {
		return nil, fmt.Errorf("cannot update Order without a valid ID")
	}

	item, err := attributevalue.MarshalMap(updatedOrder)
	if err != nil {
		return nil, err
	}

	_, err = repo.Client.PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String("Orders"),
		Item:      item,
	})
	if err != nil {
		fmt.Printf("Couldn't update item in table. Here's why: %v\n", err)
		return nil, err
	}

	return &updatedOrder, nil
}

// delete
func (repo *DynamoDBOrderRepository) DeleteOrder(orderId, productId string) error {
	if orderId == "" {
		return fmt.Errorf("cannot delete Order without a valid ID")
	}

	keyCondition := map[string]types.AttributeValue{
		"OrderId":   &types.AttributeValueMemberS{Value: orderId},
		"ProductId": &types.AttributeValueMemberS{Value: productId}, // Adjust "SortKey" based on your actual sort key attribute name
	}

	_, err := repo.Client.DeleteItem(context.Background(), &dynamodb.DeleteItemInput{
		TableName: aws.String("Orders"),
		Key:       keyCondition,
	})
	if err != nil {
		fmt.Printf("Couldn't delete item from table. Here's why: %v\n", err)
		return fmt.Errorf("failed to delete Order: %v", err)
	}

	return err
}
