package threadInterface

type ThreadInterface interface {
	notifyone()
	cutoffOccurrred() bool
	isAvailableTo(*ThreadInterface) bool
	waitfor(bool)
}
