package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	// TODO: добавить поля
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	a := strings.Split(datastring, ",")

	if len(a) != 3 {
		return errors.New("trainings - неверные данные")
	}

	a1, err := strconv.Atoi(a[0])
	if err != nil || a1 <= 0 {
		return errors.New("trainings - неверные данные шагов")
	}

	a2, err := time.ParseDuration(a[2])
	if err != nil || a2 <= 0 {
		return errors.New("trainings - неверные данные временим")
	}

	t.Steps = a1
	t.TrainingType = a[1]
	t.Duration = a2

	return err
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	distance := spentenergy.Distance(t.Steps, t.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	if t.TrainingType == "Ходьба" || t.TrainingType == "ходьба" {
		calories, err := spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)

		if err != nil {
			return "", err
		}

		str1 := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\n", t.TrainingType, t.Duration.Hours(), distance)
		str2 := fmt.Sprintf("Скорость: %.2f км/ч\nСожгли калорий: %.2f\n", meanSpeed, calories)
		return str1 + str2, err

	} else if t.TrainingType == "Бег" || t.TrainingType == "бег" {
		calories, err := spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)

		if err != nil {
			return "", err
		}

		str1 := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\n", t.TrainingType, t.Duration.Hours(), distance)
		str2 := fmt.Sprintf("Скорость: %.2f км/ч\nСожгли калорий: %.2f\n", meanSpeed, calories)
		return str1 + str2, err

	} else {
		return "", errors.New("неизвестный тип тренировки")
	}
}
