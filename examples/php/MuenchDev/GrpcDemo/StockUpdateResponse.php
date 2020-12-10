<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: protos/stock.proto

namespace MuenchDev\GrpcDemo;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>stock.StockUpdateResponse</code>
 */
class StockUpdateResponse extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>bool ack = 1;</code>
     */
    protected $ack = false;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type bool $ack
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Protos\Stock::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>bool ack = 1;</code>
     * @return bool
     */
    public function getAck()
    {
        return $this->ack;
    }

    /**
     * Generated from protobuf field <code>bool ack = 1;</code>
     * @param bool $var
     * @return $this
     */
    public function setAck($var)
    {
        GPBUtil::checkBool($var);
        $this->ack = $var;

        return $this;
    }

}

