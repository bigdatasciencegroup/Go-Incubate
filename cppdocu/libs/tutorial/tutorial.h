// #ifndef TUTORIAL_H
// #define TUTORIAL_H


/**
 * @brief Use brief, otherwise the index won't have a brief explanation.
 *
 * Detailed explanation.
 */
typedef enum BoxEnum_enum {
  BOXENUM_FIRST,  /**< Some documentation for first. */
  BOXENUM_SECOND, /**< Some documentation for second. */
  BOXENUM_ETC     /**< Etc. */
} BoxEnum;


void HelpFn2();    /*!< Internal. */

/**
 * time_now - return the current time
 *
 * Example:
 *  printf("Now is %lu seconds since epoch\n", (long)time_now().tv_sec);
 */
struct timeval time_now(void);

/**
 * time_now - return the current time
 *
 * Example:
 *  printf("Now is %lu seconds since epoch\n", (long)time_now().tv_sec);
 */
int func1(int y, double r);

/**
 *Struct1 is a basic struct representing a complex number
 */
struct struct1
{
  /**
   *ID number
   */
  int pid;
  /**
   *  Name of object
   */
	char name[]; 
};

/**
  * Struct2 is a basic struct representing a complex number
  */
struct struct2
{
  /**
    * ID number
    */
  int pid2;
  /**
   *  Name of object
   */
	char name2[]; 
};

/**
 *  Variable 1
 */
int var1;

// #endif

/*! \def MAX(a,b)
 * \brief A macro that returns the maximum of \a a and \a b.
 *
 * Details.
 */

/*! \var typedef unsigned int UINT32
 * \brief A type definition for unsigned int.
 *
 * Details.
 */

/*! \var int errno
 * \brief Contains the last error code.
 *
 * \warning Not thread safe!
 */

/*! \fn int close(int fd)
 * \brief Closes the file descriptor \a fd.
 * \param fd The descriptor to close.
 */

#define MAX(a,b) (((a) > (b)) ? (a) : (b))
typedef unsigned int UINT32;
int errno;
int close(int);