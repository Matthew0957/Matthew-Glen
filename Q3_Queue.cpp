#include <iostream>

using namespace std;

const int MAX_SIZE = 5;
int queueArr[MAX_SIZE];
int frontIdx = 0;
int rearIdx = -1;
int currentSize = 0;

bool isFull() {
    return currentSize == MAX_SIZE;
}

bool isEmpty() {
    return currentSize == 0;
}

void enqueue(int val) {
    if (isFull()) {
        cout << "Queue Overflow! Cannot enqueue " << val << ".\n";
    } else {
        rearIdx = (rearIdx + 1) % MAX_SIZE;
        queueArr[rearIdx] = val;
        currentSize++;
        cout << val << " inserted successfully.\n";
    }
}

void dequeue() {
    if (isEmpty()) {
        cout << "Queue Underflow! No elements to dequeue.\n";
    } else {
        cout << queueArr[frontIdx] << " removed successfully.\n";
        frontIdx = (frontIdx + 1) % MAX_SIZE;
        currentSize--;
    }
}

void displayFront() {
    if (isEmpty()) {
        cout << "Queue is empty.\n";
    } else {
        cout << "Front element: " << queueArr[frontIdx] << "\n";
    }
}

void display() {
    if (isEmpty()) {
        cout << "Queue: (empty)\n";
    } else {
        cout << "Queue: ";
        for (int i = 0; i < currentSize; i++) {
            int idx = (frontIdx + i) % MAX_SIZE;
            cout << queueArr[idx] << " ";
        }
        cout << "\n";
    }
}

int main() {
    int choice, value;
    do {
        cout << "\n===== QUEUE MENU =====\n";
        cout << "1. Enqueue\n2. Dequeue\n3. Front\n4. Display\n5. Exit\n";
        cout << "Choose operation: ";
        cin >> choice;

        switch (choice) {
            case 1:
                cout << "Enter value: ";
                cin >> value;
                enqueue(value);
                break;
            case 2:
                dequeue();
                break;
            case 3:
                displayFront();
                break;
            case 4:
                display();
                break;
            case 5:
                cout << "Exiting program...\n";
                break;
            default:
                cout << "Invalid choice!\n";
        }
    } while (choice != 5);

    return 0;
}
