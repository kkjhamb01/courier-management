CREATE TABLE ride_location
(
    id                   BINARY(36) DEFAULT UUID(),
    request_id           BINARY(36),
    lat                  double,
    lng                  double,
    full_name            VARCHAR(50),
    phone_number         VARCHAR(50),
    address_details      VARCHAR(50),
    courier_instructions VARCHAR(50),

    location_order       int,
    type                 ENUM ('source', 'destination', 'pickup'),

    primary key (id),
    foreign key (request_id) references request (id)
);

CREATE TABLE request
(
    id                       BINARY(36) DEFAULT UUID(),
    customer_id              BINARY(36),
    courier_id               BINARY(36),
    estimated_duration       bigint,
    estimated_distance_meter int,
    final_price              double,
    final_price_currency     VARCHAR(3),
    vehicle_type             ENUM ('bicycle', 'truck', 'first_available'),
    required_workers         int,
    human_readable_id        VARCHAR(20),
    status                   ENUM ('new', 'wait4accept', 'cancelled'),
    cancelling_reason        ENUM ('reason1', 'reason2'),
    cancelled_by             ENUM ('client', 'courier'),
    created_at               timestamp,
    updated_by               BINARY(36) DEFAULT UUID(),
    updated_at               timestamp,

    primary key (id)
);

CREATE table ride_event
(
    id              BINARY(36) DEFAULT UUID(),
    request_id      BINARY(36),
    event_type      bool,
    additional_info VARCHAR(200),
    created_at      timestamp,

    primary key (id),
    foreign key (request_id) references request (id)
)