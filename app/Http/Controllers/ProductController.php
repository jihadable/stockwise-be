<?php

namespace App\Http\Controllers;

use App\Models\Product;
use App\Utils\ResponseDefault;
use Faker\Factory as Faker;
use Illuminate\Http\Request;
use Tymon\JWTAuth\Facades\JWTAuth;

class ProductController extends Controller {
    public function store(Request $request){
        $user = JWTAuth::parseToken()->authenticate();
        $user_id = $user->id;
        $slug = Faker::create()->uuid;

        Product::create([...$request->all(), "user_id" => $user_id, "slug" => $slug]);

        return response()->json(ResponseDefault::create(202, true, "Stored product successfully"), 202);
    }

    public function delete($slug){
        $product = Product::where("slug", $slug)->first();

        if (!$product){
            return response()->json(ResponseDefault::create(404, false, "Product not found"), 404);
        }
        
        return response()->json(["slug" => $slug]);
    }

    public function update(Request $request, $slug){
        
    }
}
