<?php

namespace App\Http\Controllers;

use App\Models\Product;
use App\Utils\ResponseDefault;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Validator;
use Tymon\JWTAuth\Facades\JWTAuth;
use Illuminate\Support\Str;

class ProductController extends Controller {
    public function index(){
        $user = JWTAuth::parseToken()->authenticate();
        $products = $user->products()->orderBy("created_at", "desc")->get();

        return response()->json([
            ...ResponseDefault::create(200, true, "Berhasil mendapatkan semua produk"),
            "products" => $products->map(function($product){
                return $product->response();
            })
        ], 200);
    }

    public function store(Request $request){
        $validator = Validator::make($request->all(), [
            "name" => "required|string",
            "category" => "required|string",
            "price" => "required|integer|min:1",
            "quantity" => "required|integer|min:1",
            "description" => "required|string",
            "image" => "sometimes|file|mimes:jpg,jpeg,png|max:2048"
        ]);

        if ($validator->fails()) {
            return response()->json(ResponseDefault::create(400, false, $validator->errors()->first()), 400);
        }

        $user = JWTAuth::parseToken()->authenticate();
        $user_id = $user->id;
        $slug = (string) Str::uuid();

        if ($request->hasFile("image")){
            $imageData = file_get_contents($request->file('image')->getRealPath());
            $base64Image = base64_encode($imageData);
            $mimeType = $request->file('image')->getClientMimeType();
            $imagePath = "data:$mimeType;base64,$base64Image";
        }
        else {
            $imagePath = null;
        }

        $product = Product::create([...$request->all(), "user_id" => $user_id, "slug" => $slug, "image" => $imagePath]);

        return response()->json([
            ...ResponseDefault::create(202, true, "Berhasil menambah produk baru"),
            "product" => $product->response()
        ], 202);
    }

    public function delete($slug){
        $product = Product::where("slug", $slug)->first();

        if (!$product){
            return response()->json(ResponseDefault::create(404, false, "Produk tidak ditemukan"), 404);
        }

        $product->delete();
        
        return response()->json(ResponseDefault::create(200, true, "Berhasil menghapus produk"), 200);
    }

    public function update(Request $request, $slug){
        $validator = Validator::make($request->all(), [
            "name" => "required|string",
            "category" => "required|string",
            "price" => "required|integer|min:1",
            "quantity" => "required|integer|min:1",
            "description" => "required|string",
            "image" => "sometimes|file|mimes:jpg,jpeg,png|max:2048"
        ]);

        if ($validator->fails()) {
            return response()->json(ResponseDefault::create(400, false, $validator->errors()->first()), 400);
        }

        $product = Product::where("slug", $slug)->first();

        if (!$product){
            return response()->json(ResponseDefault::create(404, false, "Produk tidak ditemukan"), 404);
        }

        if ($request->hasFile("image")){
            $imageData = file_get_contents($request->file('image')->getRealPath());
            $base64Image = base64_encode($imageData);
            $mimeType = $request->file('image')->getClientMimeType();
            $imagePath = "data:$mimeType;base64,$base64Image";
        }
        else {
            $imagePath = $product->image;
        }

        $product->update([...$request->all(), "image" => $imagePath]);

        return response()->json([
            ...ResponseDefault::create(200, true, "Berhasil memperbarui data produk"),
            "product" => $product->response()
        ], 200);
    }
}
