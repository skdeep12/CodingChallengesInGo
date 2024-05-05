package MetroCard

type MetroCard interface {
	GetNumber() string
	GetBalance() float64
	DeductCharge(float64) (float64, error)
	Recharge(int) (float64, error)
	Equals(MetroCard) bool
}

type CardRepository struct {
	cards []MetroCard
}

func NewCardRepository() *CardRepository {
	return &CardRepository{
		cards: make([]MetroCard, 0),
	}
}

func (c *CardRepository) AddCard(m MetroCard) {
	c.cards = append(c.cards, m)
}

func (c *CardRepository) GetCard(cardNumber string) MetroCard {
	for _, card := range c.cards {
		if card.GetNumber() == cardNumber {
			return card
		}
	}
	return nil
}

func (c *CardRepository) AllCards() []MetroCard {
	return c.cards
}

type Card struct {
	Number  string
	Balance float64
}

func NewCard(number string, balance float64) *Card {
	return &Card{
		Number:  number,
		Balance: balance,
	}
}

func (c *Card) GetNumber() string {
	return c.Number
}

func (c *Card) GetBalance() float64 {
	return c.Balance
}

func (c *Card) Equals(c2 MetroCard) bool {
	if c.GetNumber() == c2.GetNumber() {
		return true
	}
	return false
}

func (c *Card) DeductCharge(charge float64) (float64, error) {
	c.Balance -= charge
	return c.Balance, nil
}

func (c *Card) Recharge(rechargeAmount int) (float64, error) {
	println(c.Number, c.Balance, "Recharge amount: ", rechargeAmount)
	c.Balance += float64(rechargeAmount)
	return c.Balance, nil
}
