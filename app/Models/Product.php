<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

class Product extends Model {
    use HasFactory;

    protected $fillable = [
        "user_id",
        "slug",
        "name",
        "category",
        "price",
        "quantity",
        "image",
        "description"
    ];

    public function response(){
        return [
            "user" => $this->user->response(),
            "slug" => $this->slug,  
            "name" => $this->name,
            "category" => $this->category,
            "price" => $this->price,
            "quantity" => $this->quantity,
            "image" => $this->image,
            "description" => $this->description,
        ];
    }

    public function user(): BelongsTo {
        return $this->belongsTo(User::class);
    }
}