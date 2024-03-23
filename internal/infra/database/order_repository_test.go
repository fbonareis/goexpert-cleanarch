package database

import (
	"database/sql"
	"testing"

	// sqlite3

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"

	"github.com/fbonareis/goexpert-cleanarch/internal/entity"
)

type OrderRepositoryTestSuite struct {
	suite.Suite
	Db *sql.DB
}

func (suite *OrderRepositoryTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	db.Exec("CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id))")
	suite.Db = db
}

func (suite *OrderRepositoryTestSuite) TearDownTest() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(OrderRepositoryTestSuite))
}

func (suite *OrderRepositoryTestSuite) TestGivenAnOrder_WhenSave_ThenShouldSaveOrder() {
	order, err := entity.NewOrder("123", 10.0, 2.0)
	suite.NoError(err)
	suite.NoError(order.CalculateFinalPrice())
	repo := NewOrderRepository(suite.Db)
	err = repo.Save(order)
	suite.NoError(err)

	var orderResult entity.Order
	err = suite.Db.QueryRow("Select id, price, tax, final_price from orders where id = ?", order.ID).
		Scan(&orderResult.ID, &orderResult.Price, &orderResult.Tax, &orderResult.FinalPrice)

	suite.NoError(err)
	suite.Equal(order.ID, orderResult.ID)
	suite.Equal(order.Price, orderResult.Price)
	suite.Equal(order.Tax, orderResult.Tax)
	suite.Equal(order.FinalPrice, orderResult.FinalPrice)
}
func (suite *OrderRepositoryTestSuite) TestGivenAnOrder_WhenList_ThenShouldListAllOrders() {
	order1, err := entity.NewOrder("1", 10.0, 2.0)
	suite.NoError(err)
	suite.NoError(order1.CalculateFinalPrice())
	_, err = suite.Db.Exec("INSERT INTO orders (id, price, tax, final_price) VALUES ($1, $2, $3, $4)", order1.ID, order1.Price, order1.Tax, order1.FinalPrice)
	suite.NoError(err)

	order2, err := entity.NewOrder("2", 50.0, 3.0)
	suite.NoError(err)
	suite.NoError(order2.CalculateFinalPrice())
	_, err = suite.Db.Exec("INSERT INTO orders (id, price, tax, final_price) VALUES ($1, $2, $3, $4)", order2.ID, order2.Price, order2.Tax, order2.FinalPrice)
	suite.NoError(err)

	repo := NewOrderRepository(suite.Db)
	ordersResult, err := repo.GetAll()
	suite.NoError(err)
	suite.Len(ordersResult, 2)
}
