// This is a temporary directory until all services are developing under the same repository
// The packages would probably need to be copied into the separate repositories when we decide to segregate services' source code
// This directory contains the code which all services would depend on, please do not make dramatic changes without discussing with all backend teams
// The memory accessed through this directory potentially could be accessed by multiple go routines, please consider this before making changes on the code within the directory
package common
