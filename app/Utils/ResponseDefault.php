<?php 

namespace App\Utils;

class ResponseDefault {
    public static function create($status, $ok, $message){
        return [
            "status" => $status,
            "ok" => $ok,
            "message" => $message
        ];
    }
}