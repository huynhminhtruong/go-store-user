# global environments
# services
SERVICE = book-service
KONG_REGISTER_SERVICE_URL = "url=http://gateway:8081"
KONG_REGISTER_SERVICE_NAME = "name=user-service"

# endpoint urls
KONG_ROUTE_PATH_GET_ONE = "paths[]=/v1/users"
KONG_ROUTE_PATH_GET_MANY = "paths[]=/v1/users/{user_id}"
KONG_ROUTE_PATH_POST = "paths[]=/v1/users"

# options
IS_STRIP_PATH = "strip_path=false"
HTTP_GET = "methods[]=GET"
HTTP_POST = "methods[]=POST"
HTTP_DELETE = "methods[]=DELETE"
HTTP_PUT = "methods[]=PUT"
HTTP_PATCH = "methods[]=PATCH"

# add service to kong
kong-register-service:
	@curl -i -X POST http://localhost:8001/services \
	--data ${KONG_REGISTER_SERVICE_NAME} \
	--data ${KONG_REGISTER_SERVICE_URL}

# add route to service
kong-register-route-get-one:
	@curl -i -X POST http://localhost:8001/services/${SERVICE}/routes \
      	--data ${KONG_ROUTE_PATH_GET_ONE} \
	--data ${HTTP_GET} \
	--data ${IS_STRIP_PATH}

# add route to service
kong-register-route-get-many:
	@curl -i -X POST http://localhost:8001/services/${SERVICE}/routes \
       	--data ${KONG_ROUTE_PATH_GET_MANY} \
	--data ${HTTP_GET} \
	--data ${IS_STRIP_PATH}

# add route to service
kong-register-route-post:
	@curl -i -X POST http://localhost:8001/services/${SERVICE}/routes \
       	--data ${KONG_ROUTE_PATH_POST} \
	--data ${HTTP_POST} \
	--data ${IS_STRIP_PATH}

# kong admin
# get list of created services
kong-admin-get-services:
	@curl -i -X GET http://localhost:8001/services

# get list of created routes of specific service
kong-admin-get-routes:
	@curl -i -X GET http://localhost:8001/services/${SERVICE}/routes