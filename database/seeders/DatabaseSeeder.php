<?php

namespace Database\Seeders;

// use Illuminate\Database\Console\Seeds\WithoutModelEvents;

use App\Models\User;
use Illuminate\Database\Seeder;
use Illuminate\Support\Facades\Hash;

class DatabaseSeeder extends Seeder {
    /**
     * Seed the application's database.
     */
    public function run(): void {
        $password = env("PRIVATE_PASSWORD");

        User::create([
            "username" => "umar",
            "email" => "umarjihad@gmail.com",
            "password" => Hash::make($password),
            "bio" => "test"
        ]);
    }
}
