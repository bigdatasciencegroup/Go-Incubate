add_executable(apps ./main.cpp)

target_link_libraries(apps lec04)
target_link_libraries(apps helloworld)
target_link_libraries(apps computepi)