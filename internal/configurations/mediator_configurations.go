package configurations

//func ConfigMediator(
//	ctx context.Context,
//	repository contracts.IOrderRepository,
//	kafkaReader *kafka.Reader,
//	kafkaWriter *kafka.Writer,) (err error) {
//
//	err = mediatr.RegisterRequestHandler[
//		*addOrderCommand.AddOrderCommand,
//		*addOrderCommand.AddOrderResponse](addOrderCommand.NewAddOrderHandler(repository, ctx))
//
//	err = mediatr.RegisterRequestHandler[
//		*getOrderQueries.GetOrderQuery,
//		*getOrderQueries.GetOrderResponse](getOrderQueries.NewGetOrderHandler(repository, ctx, kafkaReader, kafkaWriter))
//
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
