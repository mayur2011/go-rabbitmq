package main

func (app *Config) publishEventToRMQ(msg MessagePayload) error {
	if err := app.pushToQueue(msg.Name, msg.Data); err != nil {
		return err
	}

	return nil
}

func (app *Config) pushToQueue(name, msg string) error{
	emitter, err := event.NewEventEmitter(app.Rabbit)
	if err != nil{
		return err
	}
	message := MessagePayload{
		Name: name,
		Data: msg,
	}
	j, _ := json.MarshalIndent(&message, "", "\t")
	if err = emitter.Push(string(j), "log.INFO"); err != nil {
		return err
	}
	return nil	
}