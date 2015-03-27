package db
import "database/sql"

type puck struct {
    puck_id int32
    device_id string
    model string
    network_id string
    serial_number string
    status int16
    city_id int32
    space_id int32
}

func delTempPuck(berth string,addr string,db *sql.DB){
    //delete temp data

}

