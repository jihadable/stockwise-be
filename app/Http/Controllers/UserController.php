<?php

namespace App\Http\Controllers;

use App\Models\User;
use App\Utils\ResponseDefault;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Hash;
use Illuminate\Support\Facades\Validator;
use Tymon\JWTAuth\Facades\JWTAuth;

class UserController extends Controller {
    public function get_user_profile(){
        $user = JWTAuth::parseToken()->authenticate();

        return response()->json([
            ...ResponseDefault::create(200, true, "Berhasil mendapatkan data pengguna"),
            "user" => $user->response()
        ], 200);
    }

    public function register(Request $request){
        $validator = Validator::make($request->all(), [
            "username" => "required|string",
            "email" => "required|string|email|unique:users",
            "password" => "required|string|min:8",
        ]);
    
        if ($validator->fails()) {
            return response()->json(ResponseDefault::create(400, false, $validator->errors()->first()), 400);
        }

        $user = User::where("email", $request->email)->first();

        if ($user){
            return response()->json(
                ResponseDefault::create(400, false, "Email yang dimasukkan sudah terdaftar"), 400
            );
        }

        $user = User::create([
            "username" => $request->username,
            "email" => $request->email,
            "password" => Hash::make($request->password)
        ]);

        $token = JWTAuth::fromUser($user);

        return response()->json([
            ...ResponseDefault::create(202, true, "Pengguna berhasil registrasi"),
            "token" => $token,
            "user" => $user->response()
        ], 202);
    }
    
    public function login(Request $request){
        $validator = Validator::make($request->all(), [
            "email" => "required|string|email",
            "password" => "required|string|min:8",
        ]);
    
        if ($validator->fails()) {
            return response()->json(ResponseDefault::create(400, false, $validator->errors()->first()), 400);
        }

        $user = User::where("email", $request->email)->first();

        if (!$user || !Hash::check($request->password, $user->password)){
            return response()->json(ResponseDefault::create(400, false, "Email atau password salah"), 400);
        }

        $token = JWTAuth::fromUser($user);
        
        return response()->json([
            ...ResponseDefault::create(200, true, "Pengguna berhasil login"),
            "token" => $token,
            "user" => $user->response()
        ], 200);
    }

    public function update(Request $request){
        $validator = Validator::make($request->all(), [
            "username" => "required|string",
            "bio" => "required|string",
        ]);

        if ($validator->fails()) {
            return response()->json(ResponseDefault::create(400, false, $validator->errors()->first()), 400);
        }

        $user = JWTAuth::parseToken()->authenticate();

        $user->update([...$request->all()]);

        return response()->json(ResponseDefault::create(200, true, "Berhasil memperbarui data pengguna"));
    }
}