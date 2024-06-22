<?php

namespace App\Http\Controllers;

use App\Models\User;
use App\Utils\ResponseDefault;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Hash;
use Tymon\JWTAuth\Facades\JWTAuth;

class UserController extends Controller{
    public function get_user_profile(){
        $user = JWTAuth::parseToken()->authenticate();

        return response()->json([
            ...ResponseDefault::create(200, true, "Get user profile successfully"),
            "user" => $user->response()
        ], 200);
    }

    public function register(Request $request){
        $user = User::where("email", $request->email)->first();

        if ($user){
            return response()->json(
                ResponseDefault::create(400, false, "User have already registered"), 400
            );
        }

        $user = User::create([
            "username" => $request->username,
            "email" => $request->email,
            "password" => Hash::make($request->password)
        ]);

        $token = JWTAuth::fromUser($user);

        return response()->json([
            ...ResponseDefault::create(202, true, "User registered successfully"),
            "token" => $token
        ], 202);
    }
    
    public function login(Request $request){
        $user = User::where("email", $request->email)->first();

        if (!$user || !Hash::check($request->password, $user->password)){
            return response()->json(ResponseDefault::create(400, false, "Invalid email or password"), 400);
        }

        $token = JWTAuth::fromUser($user);
        
        return response()->json([
            ...ResponseDefault::create(200, true, "User logged in successfully"),
            "token" => $token
        ], 200);
    }
}
