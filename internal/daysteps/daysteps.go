package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	a := strings.Split(datastring, ",")

	if len(a) != 2 {
		return errors.New("daysteps - неверные данные")
	}

	a1, err := strconv.Atoi(a[0])
	if err != nil || a1 <= 0 {
		return errors.New("daysteps - неверные данные шагов")
	}

	a2, err := time.ParseDuration(a[1])
	if err != nil || a2 <= 0 {
		return errors.New("daysteps - неверные данные временим")
	}

	ds.Steps = a1
	ds.Duration = a2

	return err
}

func (ds DaySteps) ActionInfo() (string, error) {

	distance := spentenergy.Distance(ds.Steps, ds.Height)

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	str := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, calories)

	return str, err
}
