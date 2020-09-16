<?php

header('Content-Type: application/json');

header('Access-Control-Allow-Origin: ' . '*');
header('Access-Control-Allow-Methods: GET, PUT, POST, DELETE, OPTIONS');
header('Access-Control-Max-Age: 1000');
header('Access-Control-Allow-Headers: Content-Type, Authorization, X-Requested-With');

$JSON_DB = __DIR__.'\example-prayers.json';

$uri = $_SERVER['REQUEST_URI'];
$uri = substr($uri, 0, strpos($uri, '?'));
$method = $_SERVER['REQUEST_METHOD'];

if($uri === '/api' && $method === 'POST'){ // POST PRAYER
    $json = file_get_contents('php://input');
    $data = json_decode($json, true);

    $required_fields = ["prayer", "is_public"];
    foreach($required_fields as $field){
        if(!array_key_exists($field, $data)){
            http_response_code(400);
            exit();
        }
    }

    $jsonDB = file_get_contents($JSON_DB);
    $dataDB = json_decode($jsonDB, true);
    
    $maxID = 0;
    foreach($dataDB as $tupel){
        if($tupel['id'] > $maxID){
            $maxID = $tupel['id'];
        }
    }

    $data['id'] = $maxID + 1;
    $data['approved'] = false;
    $data['date'] = date('Y-m-d h:m:s');
    
    array_push($dataDB, $data);

    file_put_contents($JSON_DB, json_encode($dataDB));
    echo "WROTE IN DB!";
    http_response_code(201);
} else if($uri === '/api' && $method === 'GET'){ // GET PRAYERS
    
    echo file_get_contents($JSON_DB);

    http_response_code(200);
}