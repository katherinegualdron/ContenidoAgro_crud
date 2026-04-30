package config 

import(
	"database/sql" //para conexiones con sql
	"fmt"
	"log" //imprimir y manejar logs 


	_ "github.com/lib/pq" // Driver postgresSQL

)

var DB *sql.DB//Instancia global de la base de datos 

//connectDB estbalece conexion con postgresSQL 
func ConnectDB(){
	//variable para la conexion
	host:= "localhost"
	port:=5432
	user:="postgres"
	password:="postgres"
	dbname:="Agrocampo"
	schema:="contenido"

psqlInfo :=fmt.Sprintf(
	"host=%s port=%d user=%s password=%s dbname=%s search_path=%s sslmode=disable",
	host,port,user,password,dbname,schema,
)

//abrir conexion db

db, err:= sql.Open("postgres",psqlInfo)

if err!=nil {
	log.Fatal("Error al conectar:", err)
}

err =db.Ping()
if err != nil{
	log.Fatal("No se puede conectar:",err)
}

fmt.Println("conexion a base de datos exitosa!")
fmt.Println("Conectado a la db:", dbname, "Y esquema:", schema)

DB=db //Asignar a conexion global
}