<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

class UserController extends Controller{
    public function index(){

    }

    public function get_user_profile(){
        
    }

    public function register(Request $request){
        $body = $request->all();
    }
    
    public function login(Request $request){
        $body = $request->all();
    }
}
