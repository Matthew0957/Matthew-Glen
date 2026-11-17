#include <iostream>

using namespace std;

const int MAX_SIZE = 5;
int stackArr[MAX_SIZE];
int top = -1;

bool isFull() {
    return top == MAX_SIZE - 1;
}

bool isEmpty() {
    return top == -1;
}

void push(int val) {
    if (isFull()) {
        cout << "Stack Overflow! Cannot push " << val << ".\n";
    } else {
        top++;
        stackArr[top] = val;
        cout << val << " inserted successfully.\n";
    }
}

void pop() {
    if (isEmpty()) {
        cout << "Stack Underflow! No elements to pop.\n";
    } else {
        cout << stackArr[top] << " removed successfully.\n";
        top--;
    }
}

void peek() {
    if (isEmpty()) {
        cout << "Stack is empty. No top element.\n";
    } else {
        cout << "Top element: " << stackArr[top] << "\n";
    }
}

void display() {
    if (isEmpty()) {
        cout << "Stack is empty.\n";
    } else {
        cout << "Stack elements (top to bottom): ";
        for (int i = top; i >= 0; i--) {
            cout << stackArr[i] << " ";
        }
        cout << "\n";
    }
}

int main() {
    int choice, value;
    do {
        cout << "\n===== STACK MENU =====\n";
        cout << "1. Push\n2. Pop\n3. Peek\n4. Display\n5. Exit\n";
        cout << "Choose operation: ";
        cin >> choice;

        switch (choice) {
            case 1:
                cout << "Enter value: ";
                cin >> value;
                push(value);
                break;
            case 2:
                pop();
                break;
            case 3:
                peek();
                break;
            case 4:
                display();
                break;
            case 5:
                cout << "Exiting program...\n";
                break;
            default:
                cout << "Invalid choice! Please choose between 1-5.\n";
        }
    } while (choice != 5);

    return 0;
}
