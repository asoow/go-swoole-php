<?php

declare(strict_types=1);
/**
 * This file is part of Hyperf.
 *
 * @link     https://www.hyperf.io
 * @document https://hyperf.wiki
 * @contact  group@hyperf.io
 * @license  https://github.com/hyperf/hyperf/blob/master/LICENSE
 */

namespace App\Controller;

use Hyperf\Utils\ApplicationContext;
use Hyperf\DbConnection\Db;

class IndexController extends AbstractController
{
    public function index()
    {
//        $user = $this->request->input('user', 'Hyperf');
//        $method = $this->request->getMethod();

        return [
            'result' => 'hello'
        ];
    }

    public function dbs()
    {
        $data = ['title' => 'name'.self::random()];

        $id = Db::table('tests')->insertGetId($data);

        $data = Db::table('tests')->where('id', $id)->first();

        $data = Db::table('tests')->where('id',$id)->delete();

        return [
            'result' => $data
        ];
    }

    private function random()
    {
        return mt_rand(1, 99999999);
    }

    public function rds()
    {
        $container = ApplicationContext::getContainer();
        $redis = $container->get(\Hyperf\Redis\Redis::class);
        $key = 'myKey:'.$this->random();
        $redis->set($key, $this->random(), 3600);
        $result = $redis->get($key);
        $redis->del($key);
        return [
            'result' => $result
        ];
    }
}
