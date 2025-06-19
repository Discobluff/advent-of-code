#include "day.h"
#include <CUnit/CUnit.h>
#include <CUnit/Basic.h>

void test1Input(void){
    int res = 0;
    CU_ASSERT_EQUAL(part1("01/input.txt"),res);    
}

void test2Input(void){
    int res = 0;
    CU_ASSERT_EQUAL(part2("01/input.txt"),res);    
}

int main(void) {    
    if (CUE_SUCCESS != CU_initialize_registry())
        return CU_get_error();
    
    CU_pSuite suite = CU_add_suite("testsDay01", NULL, NULL);
    if (NULL == suite) {
        CU_cleanup_registry();
        return CU_get_error();
    }
    
    if ((NULL == CU_add_test(suite, "testPart1Input", test1Input)) ||
        (NULL == CU_add_test(suite, "testPart2Input", test2Input))) {
        CU_cleanup_registry();
        return CU_get_error();
    }
    
    CU_basic_set_mode(CU_BRM_VERBOSE);
    CU_basic_run_tests();
    CU_cleanup_registry();

    return CU_get_error();
}
