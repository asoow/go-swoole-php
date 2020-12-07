<?php

namespace app\controller;

use app\BaseController;
use think\facade\Cache;
use think\facade\Db;

class Index extends BaseController
{
    public function index()
    {
        return json(["result" => "hello"]);
    }

    public function rds()
    {
        $key = "keyTest:".self::rand();
        Cache::set($key, self::rand(), 3600);

        Cache::get($key);
        $res = Cache::delete($key);
        return json(["result" => $res]);;
    }

    private static function rand()
    {
        return mt_rand(1, 99999999);
    }

    public function dbs()
    {
        $data = ['title' => 'name'.self::rand()];
        $id = Db::table('tests')->insertGetId($data);

        $data = Db::table('tests')->where('id', $id)->find();

        $res = Db::table('tests')->delete($data['id']);
        return json(["result" => $res]);
    }
}
