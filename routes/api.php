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

// users route
Route::prefix("/users")->group(function(){

    // get user profile
    Route::get("/", [UserController::class, "get_user_profile"])->middleware(JwtMiddleware::class);
    
    // register
    Route::post("/register", [UserController::class, "register"]);
    
    // login
    Route::post("/login", [UserController::class, "login"]);
});

// products route
Route::prefix("/products")->group(function(){

    // get all products by user
    Route::get("/", [ProductController::class, "index"])->middleware(JwtMiddleware::class);
    
    // store product
    Route::post("/", [ProductController::class, "store"])->middleware(JwtMiddleware::class);
    
    // delete product
    Route::delete("/{slug}", [ProductController::class, "delete"])->middleware(JwtMiddleware::class);
    
    // update product
    Route::patch("/{slug}", [ProductController::class, "update"])->middleware(JwtMiddleware::class);
});