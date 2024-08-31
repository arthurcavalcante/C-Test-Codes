package main

type OperationSystem interface {

	CollectOperation() (float32, float32, string, error);

}