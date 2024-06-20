<?php

use App\Http\Controllers\UserController;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;

/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. These
| routes are loaded by the RouteServiceProvider and all of them will
| be assigned to the "api" middleware group. Make something great!
|
*/

// get user profile
Route::get("/users", [UserController::class, "get_user_profile"]);

// register
Route::get("/users", [UserController::class, "register"]);

// login
Route::get("/users", [UserController::class, "login"]);
