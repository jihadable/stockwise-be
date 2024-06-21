<?php

use App\Http\Controllers\ProductController;
use App\Http\Controllers\UserController;
use App\Http\Middleware\JwtMiddleware;
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
Route::get("/users", [UserController::class, "get_user_profile"])->middleware(JwtMiddleware::class);

// register
Route::post("/users/register", [UserController::class, "register"]);

// login
Route::post("/users/login", [UserController::class, "login"]);

// store product
Route::post("/products", [ProductController::class, "store"])->middleware(JwtMiddleware::class);

// delete product
Route::delete("/products/{slug}", [ProductController::class, "delete"])->middleware(JwtMiddleware::class);

// update product
Route::patch("/products/{slug}", [ProductController::class, "update"])->middleware(JwtMiddleware::class);