package seed

import (
	"github.com/bxcodec/faker/v3"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/baan"
	size "github.com/isd-sgcu/rnkm65-backend/src/constant/baan"
	"math/rand"
)

func (s Seed) BaanSeed1658079351033() error {
	baanSize := []size.BaanSize{size.S, size.M, size.L, size.XL, size.XXL}

	for i := 0; i < 10; i++ {
		ban := baan.Baan{
			NameTH:        faker.Word(),
			DescriptionTH: faker.Paragraph(),
			NameEN:        faker.Word(),
			DescriptionEN: faker.Paragraph(),
			Size:          baanSize[rand.Intn(len(baanSize))],
			Facebook:      faker.Word(),
			Instagram:     faker.Word(),
			Line:          faker.Word(),
		}
		err := s.db.Create(&ban).Error

		if err != nil {
			return err
		}
	}
	return nil
}
