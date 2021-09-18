### Para levantar los servicios en la consola ingresar
 >go run main.go


## endpoints

### Method: POST
* http://localhost:8080/login
 ```sh
{
	"email": "developer@gmail.com",
	"password": "password"
}
```
* http://localhost:8080/register
```sh
{
	"userID": "asalazarj",
	"user-name" : "JOSE ARMANDO SALAZAR JAUREGUI",
	"phoneID": "12313131",
	"password": "Password$12",
	"email": "developer@gmail.com"
}
```
* http://localhost:8080/supplier/addservice
```sh
{
	"servicename": "comida",
	"description": "asdad",
	"price":       24,
	"available":   true,
	"category":   "comida"
}
```
### Method: PUT

* http://localhost:8080/supplier/wallet/recharge
```sh
{
	"credits" : 6987
}
```
### Method: GET
  
* http://localhost:8080/supplier/wallet/historic
 > usuarioid: identificador del usuario generado de la base de datos.
 > pagina: paginacion del usuario
  ```sh
   params: 
   usuarioid : 61439418fce25e30f4424c55
   pagina  :  1
  ```
 
* http://localhost:8080/liberet/services
 > status: estado de los servicios
  ```sh
   params: 
   status : A 
  ```
