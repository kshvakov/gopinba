create database pinba_reports;

use pinba_reports;

create table report_by_group (
   tag_value varchar(64) default null,
   req_count int(11) default null,
   req_per_sec float default null,
   hit_count int(11) default null,
   hit_per_sec float default null,
   timer_value float default null,
   timer_median float default null,
   index_value varchar(256) default null
) engine=pinba default charset=latin1 comment='tag_info:group';

create table report_by_group_operation (
   group_value varchar(64) default null,
   operation_value varchar(64) default null,
   req_count int(11) default null,
   req_per_sec float default null,
   hit_count int(11) default null,
   hit_per_sec float default null,
   timer_value float default null,
   timer_median float default null,
   index_value varchar(256) default null
) engine=pinba default charset=latin1 comment='tag2_info:group,operation';

create table report_by_group_operation_script (
   script_name varchar(128) default null,
   group_value varchar(64) default null,
   operation_value varchar(64) default null,
   req_count int(11) default null,
   req_per_sec float default null,
   hit_count int(11) default null,
   hit_per_sec float default null,
   timer_value float default null,
   timer_median float default null,
   index_value varchar(256) default null
) engine=pinba default charset=latin1 comment='tag2_report:group,operation';


create table report_by_group_script_host (
   script_name varchar(128) default null,
   tag_value varchar(64) default null,
   req_count int(11) default null,
   req_per_sec float default null,
   hit_count int(11) default null,
   hit_per_sec float default null,
   timer_value float default null,
   hostname varchar(32) default null,
   server_name varchar(64) default null,
   timer_median float default null,
   index_value varchar(256) default null
) engine=pinba default charset=latin1 comment='tag_report2:group';


CREATE TABLE report_by_group_operation_from_to (
   script_name varchar(128) default null,
   group_value varchar(64) default null,
   operation_value varchar(64) default null,
   from_server_value varchar(64) default null,
   to_server_value varchar(64) default null,
   req_count int(11) default null,
   req_per_sec float default null,
   hit_count int(11) default null,
   hit_per_sec float default null,
   timer_value float default null,
   timer_median float default null,
   index_value varchar(256) default null
 ) engine=pinba default charset=latin1 comment='tagN_report:group,operation,from_server,to_server';
