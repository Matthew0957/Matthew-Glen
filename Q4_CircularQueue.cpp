#include <iostream>

using namespace std;

const int MAX_SIZE = 5;
int cQueue[MAX_SIZE];
int front = 0, rear = -1, count = 0;

void enqueue(int val) {
    if (count == MAX_SIZE) {
        cout << "Queue Full!\n";
    } else {
        rear = (rear + 1) % MAX_SIZE;
        cQueue[rear] = val;
        count++;
    }
}

void dequeue() {
    if (count == 0) {
        cout << "Queue Empty!\n";
    } else {
        front = (front + 1) % MAX_SIZE;
        count--;
    }
}

void display() {
    cout << "Queue: ";
    for (int i = 0; i < count; i++) {
        int idx = (front + i) % MAX_SIZE;
        cout << cQueue[idx] << " ";
    }
    cout << "\n";
}

int main() {

    enqueue(10);
    enqueue(20);
    enqueue(30);
    enqueue(40);
    enqueue(50);

    dequeue();
    dequeue();

    enqueue(60);
    enqueue(70);

    display();

    return 0;
}
