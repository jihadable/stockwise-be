<?php

namespace App\Http\Controllers;

use App\Models\User;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Hash;

class UserController extends Controller{
    public function get_user_profile(){
        
    }

    public function register(Request $request){
        $user = User::where("email", $request->email)->first();

        if ($user){
            return response()->json([
                "status" => 400,
                "ok" => false,
                "message" => "User have already registered"
            ], 400);
        }

        User::create([
            "username" => $request->username,
            "email" => $request->email,
            "password" => Hash::make($request->password),
            "image" => $request->image
        ]);

        return response()->json([
            "status" => 200,
            "ok" => true,
            "message" => "User registered successfully"
        ], 200);
    }
    
    public function login(Request $request){
        $user = User::where("email", $request->email)->first();

        if (!$user){
            return response()->json([
                "status" => 404,
                "ok" => false,
                "message" => "Invalid email or password"
            ], 404);
        }

        if (!Hash::check($request->password, $user->password)){
            return response()->json([
                "status" => 400,
                "ok" => false,
                "message" => "Invalid email or password"
            ], 400);
        }
        
        return response()->json([
            "status" => 200,
            "ok" => true,
            "message" => "User logged in successfully"
        ], 200);
    }
}
