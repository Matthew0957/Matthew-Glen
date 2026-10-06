#include <iostream>
#include <string>

using namespace std;

// ==========================================
// TASK 1: Create the Node
// ==========================================
struct Node {
    string data;
    Node* prev;
    Node* next;


    Node(string val) {
        data = val;
        prev = nullptr;
        next = nullptr;
    }
};


class DoublyLinkedList {
private:
    Node* head;
    Node* tail;

public:
    DoublyLinkedList() {
        head = nullptr;
        tail = nullptr;
    }


    void append(string val) {
        Node* newNode = new Node(val);
        if (head == nullptr) {
            head = tail = newNode;
        } else {
            tail->next = newNode;
            newNode->prev = tail;
            tail = newNode;
        }
    }

    // ==========================================
    // TASK 3: Forward Traversal
    //Which pointer is used to move forward? -> Pointer 'next'
    // ==========================================
    void traverseForward() {
        cout << "Forward: ";
        Node* current = head;
        while (current != nullptr) {
            cout << current->data;
            if (current->next != nullptr) cout << " <-> ";
            current = current->next;
        }
        cout << endl;
    }

    // ==========================================
    // TASK 4: Backward Traversal
    //Which pointer is used to move backward? -> Pointer 'prev'
    // ==========================================
    void traverseBackward() {
        cout << "Backward: ";
        Node* current = tail;
        while (current != nullptr) {
            cout << current->data;
            if (current->prev != nullptr) cout << " <-> ";
            current = current->prev;
        }
        cout << endl;
    }

    // ==========================================
    // TASK 5: Insert in the Middle
    // ==========================================
    void insertAfter(string targetVal, string newVal) {
        Node* current = head;
        while (current != nullptr && current->data != targetVal) {
            current = current->next;
        }

        if (current == nullptr) {
            cout << "Target " << targetVal << " tidak ditemukan!" << endl;
            return;
        }

        Node* newNode = new Node(newVal);
        newNode->next = current->next;
        newNode->prev = current;

        if (current->next != nullptr) {
            current->next->prev = newNode;
        } else {
            tail = newNode;
        }
        current->next = newNode;
    }

    // ==========================================
    // TASK 6: Delete a Node
    // ==========================================
    void deleteNode(string targetVal) {
        Node* current = head;
        while (current != nullptr && current->data != targetVal) {
            current = current->next;
        }

        if (current == nullptr) {
            cout << "Target " << targetVal << " tidak ditemukan untuk dihapus!" << endl;
            return;
        }


        if (current == head) {
            head = head->next;
            if (head != nullptr) {
                head->prev = nullptr;
            } else {
                tail = nullptr;
            }
        }

        else if (current == tail) {
            tail = tail->prev;
            tail->next = nullptr;
        }

        else {
            current->prev->next = current->next;
            current->next->prev = current->prev;
        }

        delete current;
    }
};

int main() {
    cout << "=== DOUBLY LINKED LIST MINI LAB ==-\n" << endl;

    DoublyLinkedList dll;

    // ==========================================
    // TASK 2: Build a List (Create at least 5 nodes)
    // ==========================================
    cout << "--- TASK 2: Build a List ---" << endl;
    dll.append("Song A");
    dll.append("Song B");
    dll.append("Song C");
    dll.append("Song D");
    dll.append("Song E");


    dll.traverseForward();
    cout << endl;

    // ==========================================
    // TASK 3: Forward Traversal Execution
    // ==========================================
    cout << "--- TASK 3: Forward Traversal ---" << endl;
    dll.traverseForward();
    cout << "Question Answer: Pointer yang digunakan untuk bergerak maju adalah 'next'." << endl;
    cout << endl;

    // ==========================================
    // TASK 4: Backward Traversal Execution
    // ==========================================
    cout << "--- TASK 4: Backward Traversal ---" << endl;
    dll.traverseBackward();
    cout << "Question Answer: Pointer yang digunakan untuk bergerak mundur adalah 'prev'." << endl;
    cout << endl;

    // ==========================================
    // TASK 5: Insert in the Middle
    // ==========================================
    cout << "--- TASK 5: Insert in the Middle ---" << endl;
    cout << "Before: ";
    dll.traverseForward();

    dll.insertAfter("Song B", "Song X");

    cout << "After  : ";
    dll.traverseForward();
    cout << endl;

    // ==========================================
    // TASK 6: Delete a Node (Delete Song C)
    // ==========================================
    cout << "--- TASK 6: Delete a Node ---" << endl;
    cout << "Before: ";
    dll.traverseForward();

    dll.deleteNode("Song C");

    cout << "After  : ";
    dll.traverseForward();
    cout << "Penjelasan pointer: Ketika Song C dihapus, pointer 'next' dari Song X menunjuk ke Song D, "
         << "dan pointer 'prev' dari Song D menunjuk balik ke Song X, sehingga melewati node C." << endl;

    return 0;
}
