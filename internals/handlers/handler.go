package handlers

// type Handler struct {
// 	handler *amqp.AmqpHandler
// 	srv     ports.GenerateService
// }

// func NewHandler(srv ports.GenerateService, handler *amqp.AmqpHandler) *Handler {
// 	return &Handler{
// 		srv:     srv,
// 		handler: handler,
// 	}
// }

// func (handler *Handler) GetRecipes(msgs <-chan myrabbit.Delivery, pubCh myrabbit.Channel, subCh myrabbit.Channel) {
// 	for msg := range msgs {
// 		func() {
// 			req := &domain.Request{}
// 			defer msg.Ack(false)
// 			err := json.Unmarshal(msg.Body(), &req)
// 			if err != nil {
// 				resp := &models.Response{
// 					StatusCode: 400,
// 				}
// 				pubCh.Respond(msg, resp)
// 				return
// 			}
// 			recipes, err := handler.srv.GetRecipesFromJson("recipes.json")
// 			if err != nil {
// 				resp := &models.Response{
// 					StatusCode: 500,
// 				}
// 				pubCh.Respond(msg, resp)
// 				return
// 			}

// 			pubCh.Respond(msg, recipes)
// 		}()
// 	}
// }

// func (handler *Handler) GenerateRecipes(msgs <-chan myrabbit.Delivery, pubCh myrabbit.Channel, subCh myrabbit.Channel) {
// 	for msg := range msgs {
// 		func() {
// 			req := &domain.Request{}
// 			defer msg.Ack(false)
// 			err := json.Unmarshal(msg.Body(), &req)
// 			if err != nil {
// 				resp := &models.Response{
// 					StatusCode: 400,
// 				}
// 				pubCh.Respond(msg, resp)
// 				return
// 			}
// 			recipes, err := handler.srv.GenerateRecipes()
// 			if err != nil {
// 				resp := &models.Response{
// 					StatusCode: 500,
// 				}
// 				pubCh.Respond(msg, resp)
// 				return
// 			}

// 			pubCh.Respond(msg, recipes)
// 		}()
// 	}
// }

// func (handler *Handler) InitServer() {
// 	handler.handler.RegisterConsumer("data-generator.generate-recipes", handler.GenerateRecipes)
// }
