// This is the worker that is responsible for scheduling the events
// This scheduler is optimized for the delivery of the events in
// increasing magnitude of the sequence numbers. 
package worker

import (
	"container/heap"
	"../event"
)

// A min-heap is being used to store the events so that adding an event and extracting the 
// event that has the minimum sequence number occurs in O(log(n))
type EventHeap []event.Event

func (events EventHeap) Len() int { 
	return len(events)
}

func (events EventHeap) Less(i, j int) bool { 
	return events[i].GetSequence() < events[j].GetSequence() 
}

func (events EventHeap) Swap(i, j int) { 
	events[i], events[j] = events[j], events[i]
}

func (events *EventHeap) Push(e interface{}) {
	*events = append(*events, e.(event.Event))
}

func (events *EventHeap) Pop() interface{} {
	old := *events
	n := len(old)
	x := old[n-1]
	*events = old[0 : n-1]
	return x
}

// This is the class that is used to represent an event scheduler
// eventGenerator - this is the object that generates the events that this class will schedule
// eventHeap - this is the data structure that stores the events that have been generated but are yet to be dilivered
// maxBuffSize - this is the maximum number of events that can be stored on the machine. 
//				If the buffer becomes full, the events will no longer be send in the right order.
//				However, if the buffer never becomes full, then all events will be send in the right order.
type EventScheduler struct {
	eventGenerator *EventGenerator
	eventHeap *EventHeap
	maxBuffSize int
}

// This is the constructor for a new event scheduler
func NewEventScheduler(eventGenerator *EventGenerator, maxBuffSize int) *EventScheduler {
	eventScheduler := new(EventScheduler)
	eventScheduler.eventGenerator = eventGenerator

	eventScheduler.maxBuffSize = maxBuffSize

	eventHeap := &EventHeap{}
	eventScheduler.eventHeap = eventHeap
	heap.Init(eventScheduler.eventHeap)
	
	return eventScheduler
}

// Method for getting the event with sequence number "expected"
// The buffer is 
func (scheduler EventScheduler) NextEvent(expected int) (event.Event, bool) {
	for {
		if scheduler.eventHeap.Len() > 0 {
			min := heap.Pop(scheduler.eventHeap).(event.Event)
			if min.GetSequence() == expected {
				return min, true
			} else if (scheduler.eventHeap.Len() == scheduler.maxBuffSize) {
				// if the event buffer is full, just send the event that has the smallest sequence number
				// whether it is the one that is expected or not.
				return min, false
			} else {
				heap.Push(scheduler.eventHeap, min)
			}
		}
		e := scheduler.eventGenerator.NextEvent()
		if e != nil {
			heap.Push(scheduler.eventHeap, e)
		}
	}
} 