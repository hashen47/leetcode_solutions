#include <stdio.h>
#include <stdlib.h>
#include <stdarg.h>

// first have to create linkedlist constructor using integer array 
// then create both problem and solution using this constructor
// then we can compare problem linked list answer and real linked list solution

typedef struct ListNode ListNode;
struct ListNode {
    int val;
    struct ListNode *next;
};

// array
int *array_init(int el_count, ...);

// linked list
ListNode *linked_list_init(int *arr, int el_count);
void linked_list_free(ListNode *head);

// problem
ListNode *odd_even_list(ListNode *head);

typedef struct {
    int *problem;
    int *solution;
    int el_count; 
} Testcase;

int main() {
    Testcase testcases[] = {
        {
            // .problem = {1, 2, 3, 4, 5},
            .problem = array_init(5, 1, 2, 3, 4, 5),
            .solution = array_init(5, 1, 3, 5, 2, 4),
            .el_count = 5,
        },
        {
            .problem = array_init(10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10),
            .solution = array_init(10, 1, 3, 5, 7, 9, 2, 4, 6, 8, 10),
            .el_count = 10,
        },
    };
    int testcase_count = sizeof(testcases)/sizeof(Testcase);

    int fail_count = 0;
    for (int i = 0; i < testcase_count; i++) {
        Testcase tc = testcases[i];

        // initialize linked lists
        ListNode *problem_list = linked_list_init(tc.problem, tc.el_count);

        // solve using algorithm
        problem_list = odd_even_list(problem_list);

        // check generated solution is right
        int j = 0;
        ListNode *temp = problem_list;
        for (;;) {
            if (j >= tc.el_count) {
                fail_count++;
                fprintf(stderr, "(%d) TESTCASE FAIL: index out of range (top)\n", i);
                break;
            }
            if (temp->val != tc.solution[j]) {
                fail_count++;
                fprintf(stderr, "(%d) TESTCASE FAIL: values different: expect=%d, found=%d\n", i, tc.solution[j], temp->val);
                break;
            }
            temp = temp->next;
            j++;

            if (temp == NULL) {
                if (j != tc.el_count) {
                    fail_count++;
                    fprintf(stderr, "(%d) TESTCASE FAIL: index out of range (bottom)\n", i);
                }
                break;
            }
        }

        // free linked lists
        linked_list_free(problem_list);
    }

    if (fail_count == 0) {
        printf("ALL TESTCASES ARE PASSED\n");
    }
}

// element count can be [0, 10**4]
ListNode *odd_even_list(ListNode *head) {
    // no elements
    if (head == NULL)
        return NULL;

    // only one element
    if (head->next == NULL)
        return head;

    ListNode *odd_tail = head;
    ListNode *even_head = head->next;
    ListNode *even_tail = even_head;
    ListNode *temp = head->next->next;

    for (;;) {
        if (temp == NULL)
            break;

        odd_tail->next = temp;
        odd_tail = odd_tail->next;

        if (temp->next == NULL)
            break;

        even_tail->next = temp->next;
        even_tail = even_tail->next;
        temp = temp->next->next;
    }

    even_tail->next = NULL;
    odd_tail->next = even_head;
    return head;
}

ListNode *linked_list_init(int *arr, int el_count) {
    if (arr == NULL || el_count == 0)
        return NULL;

    ListNode *head = malloc(sizeof(ListNode));
    head->val = arr[0];

    ListNode *temp = head;
    for (int i = 1; i < el_count; i++) {
        temp->next = malloc(sizeof(ListNode));
        temp->next->val = arr[i];
        temp->next->next = NULL;
        temp = temp->next;
    }

    return head;
}

void linked_list_free(ListNode *head) {
    if (head == NULL)
        return;

    while (head != NULL) {
        ListNode *temp = head;
        head = head->next;
        free(temp);
    }
}

int *array_init(int el_count, ...) {
    int *arr = malloc(sizeof(int) * el_count);

    va_list args;
    va_start(args, el_count);

    for (int i = 0; i < el_count; i++) {
        arr[i] = va_arg(args, int);
    }

    va_end(args);
    return arr;
}

