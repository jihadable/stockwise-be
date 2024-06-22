<?php

namespace App\Http\Controllers;

use App\Models\Product;
use App\Utils\ResponseDefault;
use Faker\Factory as Faker;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Storage;
use Tymon\JWTAuth\Facades\JWTAuth;

class ProductController extends Controller {
    public function index(){
        $user = JWTAuth::parseToken()->authenticate();
        $products = $user->products()->orderBy("created_at", "desc")->paginate(10);

        return response()->json([
            ...ResponseDefault::create(200, true, "Get all products by user successfully"),
            "products" => $products->map(function($product){
                return $product->response();
            })
        ], 200);
    }

    public function store(Request $request){
        $user = JWTAuth::parseToken()->authenticate();
        $user_id = $user->id;
        $slug = Faker::create()->uuid;

        if ($request->hasFile("image")){
            $imagePath = $request->file("image")->store("images");
        }
        else {
            $imagePath = null;
        }

        Product::create([...$request->all(), "user_id" => $user_id, "slug" => $slug, "image" => $imagePath]);

        return response()->json(ResponseDefault::create(202, true, "Stored product successfully"), 202);
    }

    public function delete($slug){
        $product = Product::where("slug", $slug)->first();

        if (!$product){
            return response()->json(ResponseDefault::create(404, false, "Product not found"), 404);
        }

        if ($product->image){
            Storage::delete($product->image);
        }

        $product->delete();
        
        return response()->json(ResponseDefault::create(200, true, "Deleted product successfully"));
    }

    public function update(Request $request, $slug){
        $product = Product::where("slug", $slug)->first();

        if (!$product){
            return response()->json(ResponseDefault::create(404, false, "Product not found"), 404);
        }

        if ($request->hasFile("image")){
            if ($product->image){
                Storage::delete($product->image);
            }

            $imagePath = $request->file("image")->store("images");
        }
        else {
            $imagePath = $product->image;
        }

        $product->update([...$request->all(), "image" => $imagePath]);

        return response()->json(ResponseDefault::create(200, true, "Updated product successfully"));
    }
}
