var sql = '';

$.post('/address/getProvinces.action', function(ret) {
    for ( var provinceId in ret) {
        sql += "insert into tb_province (id, name) values ('" + provinceId + "', '" + ret[provinceId] + "');\r\n";

        (function(provinceId){
            $.post('/address/getCitys.action', {
                provinceId : provinceId
            }, function(ret) {
                for ( var cityId in ret) {
                    sql += "insert into tb_city (id, province_id, name) values ('" + cityId + "', '" + provinceId + "', '" + ret[cityId] + "');\r\n";

                    (function(cityId){
                        $.post('/address/getCountys.action', {
                            cityId : cityId
                        }, function(ret) {
                            for ( var countyId in ret) {
                                sql += "insert into tb_county (id, province_id, city_id, name) values ('" + countyId + "', '" + provinceId + "', '" + cityId + "', '" + ret[countyId] + "');\r\n";
                            }
                        });
                    })(cityId);
                }
            });        
        })(provinceId);
    }
});
