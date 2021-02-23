-- +migrate Up
create table events (
    id serial primary key,
    name varchar(255) not null,
    slug varchar(255) not null,
    description text not null,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now()
);

create table speakers (
    id serial primary key,
    name varchar(255) not null,
    bio text,
    picture_url varchar(255),
    social_networks jsonb,
    event_id int not null,

    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),

    constraint fk_speaker_event foreign key(event_id) references events(id)
);

create table talks (
    id serial primary key,
    title varchar(255) not null,
    description text not null,
    level varchar(50) not null,
    tags varchar(512),
    duration varchar(100) not null,
    event_id int not null,

    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),

    constraint fk_talk_even foreign key(event_id) references events(id)
);

create table attendees (
    id serial primary key,
    name varchar(255) not null,
    username varchar(255) not null,
    password varchar(255),
    email varchar(255) not null,
    authentication_type varchar(100) default('userpass'),
    social_networks jsonb not null default '{}'::jsonb,

    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now()
);

create table stages (
    id serial primary key,
    name varchar(255) not null,
    description text,
    event_id int,

    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),

    constraint fk_stage_event foreign key(event_id) references events(id)
);

create table notes (
    id serial primary key,
    content text not null,
    talk_id int not null,

    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),

    constraint fk_note_talk foreign key(talk_id) references talks(id)
);

create table comments (
    id serial primary key,
    content text not null,
    attendee_id int not null,
    talk_id int not null,

    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),

    constraint fk_comment_attendee foreign key(attendee_id) references attendees(id),
    constraint fk_comment_talk foreign key(talk_id) references talks(id)
);

create table speaker_table (
    id serial primary key,
    speaker_id int not null,
    talk_id int not null,

    constraint fk_spk_tlk_speaker foreign key(speaker_id) references speakers(id),
    constraint fk_spk_tlk_talk foreign key(talk_id) references talks(id)
);

create table attendee_event (
    id serial primary key,
    attendee_id int not null,
    event_id int not null,

    constraint fk_att_evt_attendee foreign key(attendee_id) references attendees(id),
    constraint fk_att_evt_event foreign key(event_id) references events(id)
);

-- +migrate Down
drop table if exists attendee_event;
drop table if exists speaker_table;
drop table if exists comments;
drop table if exists notes;
drop table if exists stages;
drop table if exists attendees;
drop table if exists talks;
drop table if exists speakers;
drop table if exists events;
