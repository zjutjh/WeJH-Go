#include <stdarg.h>
#include <stdbool.h>
#include <stdint.h>
#include <stdlib.h>

typedef struct ele_result {
  float total_surplus;
  float total_amount;
  float surplus;
  float surplus_amount;
  float subsidy;
  float subsidy_amount;
  char *display_room_name;
  char *room_status;
} ele_result;

typedef struct login_handle {
  char *phone_num;
  char *device_id;
} login_handle;

/**
 * Security token result
 * -----------
 * - `token: *mut c_char`: token c-string
 * - `level: c_int`: security level, 0: no captcha required, 1: captcha required
 */
typedef struct security_token_result {
  int level;
  char *token;
} security_token_result;

/**
 * Login result
 * -----------
 * - `bind_card_status: c_int`: 0: not bind, 1: bound
 */
typedef struct login_result {
  char *uid;
  char *token;
  char *device_id;
  int bind_card_status;
} login_result;

/**
 * Authorization -- C Bind
 * ----------
 * # Inputs
 * - `uid: *const c_char`: uid c-string, UTF-8
 * # Returns
 * - `*mut c_char`: session c-string, UTF-8. Return nullptr on error.
 */
char *auth(const char *uid);

/**
 * Query electricity -- C Bind
 * -----------
 * After calling this function,
 * the caller is responsible for using `free_ele_result` to deallocate the memory.
 *
 * # Inputs
 * - `session: *const c_char`: session c-string
 * - `result: *mut *mut ele_result`: second-level pointer for return pointer of `ele_result` struct
 *
 * # Returns
 * - `c_int`: 0 on success, otherwise error code
 * - `result`: ele_result
 *
 * # Errors (status codes)
 * - `201`: Auth expired
 * - `202`: No bind info
 * - `101`: Other error
 */
int query_ele(const char *session, struct ele_result **result);

/**
 * Free ele_result
 * -----------
 * Deallocate the struct to avoid memory leak.
 */
void free_ele_result(struct ele_result *p);

/**
 * Generate random device id -- C Bind
 * -----------
 * Generate random device into struct `login_handle`.
 *
 * The caller is responsible for using `free_c_string` to deallocate
 * the string in the structure.
 *
 * # Inputs
 * - `handle: *const login_handle`: Pointer of Login handle
 */
void gen_device_id(struct login_handle *handler);

/**
 * Get security token -- C Bind
 * -----------
 * Using filled handle to query security token.
 * # Inputs
 * - `handle: *const login_handle`: Pointer of Login handle
 * - `result: *mut *mut security_token_result`: second-level pointer for return pointer of `security_token_result` struct
 *
 * # Returns
 * - `c_int`: 0 on success, otherwise error code
 * - `result`: security_token_result
 *
 * # Errors
 * - `203`: Initialize login handler failed
 * - `101: Other errors
 */
int get_security_token(const struct login_handle *handle,
                       struct security_token_result **result);

/**
 * Free security_token_result
 * -----------
 * Deallocate the struct to avoid memory leak.
 */
void free_security_token_result(struct security_token_result *p);

/**
 * Free c_string
 * -----------
 * Deallocate c_string to avoid memory leak.
 */
void free_c_string(char *c_string);

/**
 * Get captcha image -- C Bind
 * -----------
 * Get captcha image in base64 format.
 *
 * Like `data:image/jpeg;base64,xxxxxxxxxxxxxx`
 *
 * # Inputs
 * - `handle: *const login_handle`: Pointer of Login handle
 * - `security_token: *const c_char`: security token
 * - `result: *mut *mut c_char`: second-level pointer for result
 *
 * # Returns
 * - `result: *mut *mut c_char`: captcha image in base64
 *
 * # Errors
 * - `207`: Get captcha image failed
 * - `101: Other errors
 *
 * # Usage
 *
 */
int get_captcha_image(const struct login_handle *handle, const char *security_token, char **result);

/**
 * Send SMS verification code -- C Bind
 * -----------
 * # Inputs
 * - `handle: *const login_handle`: Pointer of Login handle
 * - `security_token: *const c_char`: c-string of security token
 * - `captcha: *const c_char`: c-string of captcha.
 * If captcha input `NULL`, it means no captcha is required.
 * # Returns
 * - `c_int`: `0` on success, `1` on user is not exist(registered), otherwise error code
 * # Errors
 * - `203`: Initialize login handler failed
 * - `204`: Bad phone number
 * - `205`: Limit of SMS verification code sent
 * - `101: Other errors
 */
int send_verification_code(const struct login_handle *handle,
                           const char *security_token,
                           const char *captcha);

/**
 * Do login with verification code -- C Bind
 * -----------
 * Do login to get uid, token(APP), device id, and bind card status.
 * # Inputs
 * - `handle: *const login_handle`: Pointer of Login handle
 * - `code: *const c_char`: c-string of verification code
 * - `result: *mut *mut login_result`: second-level pointer for return pointer of `login_result` struct
 * # Returns
 * - `c_int`: 0 on success, otherwise error code
 * # Errors
 * - `203`: Initialize login handler failed
 * - `206`: Bad(Wrong) verification code
 * - `101: Other errors
 */
int do_login(const struct login_handle *handle,
             const char *code,
             struct login_result **result);

/**
 * Free login_result
 * ---------
 * Deallocate the struct to avoid memory leak.
 */
void free_login_result(struct login_result *p);