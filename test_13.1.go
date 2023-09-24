// Автоматическое тестирование команд для танков:
// - проверка, что танк правильно отвечает на валидные команды
// - проверка, что танк не отвечает на недопустимые команды
// - проверка на несуществующие или неизвестные команды

package main

import (
  "errors"
  "testing"
)

// Тест для проверки правильности выполнения валидной команды
func TestExecuteValidCommand(t *testing.T) {
  tank := &Tank{ID: "1", PlayerID: "player_1"}

  command := &Command{
    ID:             "1",
    Action:         "StartMove",
    InitialVelocity: 2,
    PlayerID:       "player_1",
  }

  err := tank.ExecuteCommand(command)

  if err != nil {
    t.Errorf("Expected no error but got %v", err)
  }

  if tank.Velocity != 2 {
    t.Errorf("Expected velocity 2 but got %v", tank.Velocity)
  }
}

// Тест для проверки ответа на недопустимую команду
func TestExecuteInvalidCommand(t *testing.T) {
  tank := &Tank{ID: "1", PlayerID: "player_1"}

  command := &Command{
    ID:             "1",
    Action:         "StartMove",
    InitialVelocity: 2,
    PlayerID:       "player_2",
  }

  err := tank.ExecuteCommand(command)

  expectedError := errors.New("unauthorized: this player cannot control this tank")

  if err == nil || err.Error() != expectedError.Error() {
    t.Errorf("Expected error '%v' but got '%v'", expectedError, err)
  }
}

// Тест для проверки на несуществующую команду
func TestExecuteUnknownCommand(t *testing.T) {
  tank := &Tank{ID: "1", PlayerID: "player_1"}

  command := &Command{
    ID:             "1",
    Action:         "Fly",
    InitialVelocity: 2,
    PlayerID:       "player_1",
  }

  err := tank.ExecuteCommand(command)

  expectedError := errors.New("unknown action")

  if err == nil || err.Error() != expectedError.Error() {
    t.Errorf("Expected error '%v' but got '%v'", expectedError, err)
  }
}
