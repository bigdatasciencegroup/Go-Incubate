add_library (tutorial "tutorial.cpp" "tutorial.h")
target_include_directories(tutorial PUBLIC .)

# Should install library properly here, but I'm focusing on just the docs