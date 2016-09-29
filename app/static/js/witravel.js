$(function() {

    Vue.config.delimiters = [ '[[', ']]' ];
    var path = $('#path').val();

    Vue.filter('meetingDate', function(value) {
        var d = value.split(' ')[0].split('-');
        var now = new Date();

        if (d[0] === now.getFullYear() + '') {
            return d[1] + '月' + d[2] + '日';
        }

        return d[0] + '年' + d[1] + '月' + d[2] + '日';
    })

    var travellist = Vue.extend({
        template : $('#travellist_view').html(),
        data : function() {
            return {
                travellist : []
            };
        },
        methods : {

        },
        ready : function() {

            var listScroll = new IScroll('#app .scroll_list', {
                probeType : 2,
                click : true,
                mouseWheel : true,
                momentum : true
            });

            listScroll.loadFinish = function() {
                this.pullUpStatus = 'hide';
                $('.pull_up_tips').removeClass('pull_status_loading').removeClass('pull_status_pulling').addClass('pull_status_hide');

                this.refresh();
            };

            // load data
            var me = this;

            var loadData = function(obj) {
                $.post(path + '/weapp/travellist', {
                    date : '',
                    destination : '',
                    meetingTime : obj.meetingTime,
                    type : obj.type, // prev, next
                    lastId : obj.lastId
                }, function(ret) {

                    if (ret.success) {

                        for ( var i in ret.data) {
                            me.travellist.push(ret.data[i]);
                        }

                        me.$nextTick(function() {
                            listScroll.refresh();
                            listScroll.loadFinish();
                        });
                    }

                });
            }

            // bind the iscroll

            // 上拉提示框状态，有hide, pulling, loading
            listScroll.pullUpStatus = 'hide';

            // 是否需要去加载更多
            listScroll.pullUpLoad = false;

            listScroll.on('scroll', function() {

                // console.log("y is :" + this.y + ",max y is :" +
                // this.maxScrollY);

                if (this.y < this.maxScrollY) {

                    if (this.pullUpStatus === 'hide') {
                        this.pullUpStatus = 'pulling';
                        $('.pull_up_tips').removeClass('pull_status_hide').removeClass('pull_status_loading').addClass('pull_status_pulling');
                    }

                    // 提示上拉加载更多
                    if (this.pullUpStatus === 'pulling') {
                        var percent = (this.maxScrollY - this.y) / 100.0;
                        percent = percent > 1 ? 1.0 : percent;
                        $('.pull_up_tips').css('opacity', percent);
                        $('.pull_sbar_in').css('width', (percent * 96) + '%');

                        if (this.y < (this.maxScrollY - 100)) {
                            // 上拉加载更多
                            this.pullUpLoad = true;
                        }
                    }
                }
            });

            listScroll.on('scrollEnd', function() {

                // console.log("scroll end. y is :" + this.y + ",max y is :" +
                // this.maxScrollY);

                // 不是正在加载，才进行加载
                if (this.pullUpStatus !== 'loading') {

                    if (this.pullUpLoad) {
                        this.pullUpStatus = 'loading';
                        $('.pull_up_tips').removeClass('pull_status_hide').removeClass('pull_status_pulling').addClass('pull_status_loading');
                        this.refresh();
                        this.scrollTo(0, this.maxScrollY, 1000, IScroll.utils.ease.elastic);

                        // do ajax
                        if (me.travellist.length > 0) {
                            var lastObj = me.travellist[me.travellist.length - 1]; 
                            loadData({
                                type : 'next',
                                lastId : lastObj['id'],
                                meetingTime : lastObj['meetingTime']
                            });
                        }
                    } else {
                        this.pullUpStatus = 'hide';
                        $('.pull_up_tips').removeClass('pull_status_loading').removeClass('pull_status_pulling').addClass('pull_status_hide');
                    }

                    this.pullUpLoad = false;
                }
            });

            loadData({
                type : 'next',
                lastId : '9223372036854775807',
                meetingTime : '2030-01-01 11:11:11'
            });
        }
    });

    var tourguide = Vue.extend({
        template : '<p>This is tourguide!</p>'
    });

    var edittravel = Vue.extend({
        template : '#edittravel_view',
        data : function() {
            return {
                title : '新建行程',
                travel : {
                    id : {
                        rule : {},
                        validate : true,
                        type : 'integer',
                        value : 0
                    },
                    destination : {
                        rule : '',
                        validate : true,
                        value : ''
                    },
                    meetingTime : {
                        validate : true,
                        value : ''
                    },
                    meetingCountry : {
                        validate : true,
                        type : 'integer',
                        value : ''
                    },
                    meetingProvince : {
                        validate : true,
                        type : 'integer',
                        value : ''
                    },
                    meetingCity : {
                        validate : true,
                        type : 'integer',
                        value : ''
                    },
                    meetingPlace : {
                        validate : true,
                        value : ''
                    },
                    returnDate : {
                        validate : true,
                        value : ''
                    },
                    description : {
                        validate : true,
                        value : ''
                    },
                    sponsorMobile : {
                        validate : true,
                        value : ''
                    },
                    sponsorWechat : {
                        validate : true,
                        value : ''
                    },
                    budget : {
                        validate : true,
                        type : 'integer',
                        value : 0
                    }
                },
                captchaId : '',
                captchaCode : '',
                captchaCodeValidate : false,
                countrys : [],
                provinces : [],
                citys : []
            };
        },
        methods : {
            // 加载省份
            onCountryChange : function(event) {
                var me = this;
                var countryId = $(event.target).val();
                me.provinces = [];
                me.citys = [];
                if (countryId) {
                    $.get(path + '/pub/getprovinces/' + countryId, function(ret) {
                        if (ret.success) {
                            me.provinces.push({
                                id : '',
                                name : '请选择'
                            });
                            for ( var i in ret.data) {
                                me.provinces.push(ret.data[i]);
                            }
                            me.travel.meetingProvince.value = '';
                            me.travel.meetingCity.value = '';
                        }
                    });
                }
            },

            // 加载城市
            onProvinceChange : function(event) {
                var me = this;
                var provinceId = $(event.target).val();
                me.citys = [];
                if (provinceId) {
                    $.get(path + '/pub/getcitys/' + provinceId, function(ret) {
                        if (ret.success) {
                            me.citys.push({
                                id : '',
                                name : '请选择'
                            });
                            for ( var i in ret.data) {
                                me.citys.push(ret.data[i]);
                            }
                            me.travel.meetingCity.value = '';
                        }
                    });
                }
            },

            // 重新加载验证码
            reloadCaptcha : function() {
                var me = this;

                $.get(path + '/pub/getcaptchaid', function(ret) {
                    if (ret.success) {
                        me.captchaId = ret.data;
                        me.captchaCodeValidate = false;
                        me.captchaCode = '';
                    }
                })
            },

            // 校验验证码
            validateCaptcha : function(event) {
                var me = this;
                if (me.captchaCode.length == 4) {
                    $.get(path + '/pub/validatecaptcha/' + me.captchaCode + '/' + me.captchaId, function(ret) {
                        if (ret.success) {
                            if (!ret.data.result) {
                                me.captchaId = ret.data.captchaId;
                                me.captchaCodeValidate = false;
                                me.captchaCode = '';
                            } else {
                                me.captchaCodeValidate = true;
                            }
                        }
                    });
                }
            },

            // 发布行程
            saveData : function() {
                var me = this;

                // 验证码是否正确
                if (!me.captchaCodeValidate) {
                    return;
                }

                var travel = {};
                for ( var i in me.travel) {
                    if (me.travel[i].value != undefined) {
                        var type = me.travel[i].type;
                        if (type === 'integer') {
                            travel[i] = parseInt(me.travel[i].value);
                        } else if (type === 'float') {
                            travel[i] = parseFloat(me.travel[i].value);
                        } else {
                            travel[i] = me.travel[i].value;
                        }
                    }
                    me.travel[i].validate = true;
                }

                console.debug(travel);

                $.post(path + '/weapp/travelvalidate', JSON.stringify(travel), function(ret) {
                    console.debug(ret);
                    if (ret.success) {
                        if (ret.data.validate) {
                            // 通过验证，提交表单
                            travel['captchaCode'] = me.captchaCode;
                            travel['captchaId'] = me.captchaId;

                            $.post(path + '/weapp/travelsave', JSON.stringify(travel), function(ret) {
                                if (ret.success) {
                                    me.$router.go('/travellist')
                                } else {

                                }
                            });

                        } else {
                            // 不通过验证
                            for ( var i in ret.data.msg) {
                                me.travel[i].validate = false;
                            }
                        }
                    }
                });
            }
        },
        ready : function() {
            // 不显示导航条
            this.$parent.showTabbar = false;
            var me = this;

            $.get(path + '/pub/getcountrys', function(ret) {
                if (ret.success) {
                    me.countrys.push({
                        id : '',
                        name : '请选择'
                    });
                    for ( var i in ret.data) {
                        me.countrys.push(ret.data[i]);
                    }
                }
            });

            $.get(path + '/pub/getcaptchaid', function(ret) {
                if (ret.success) {
                    me.captchaId = ret.data;
                }
            })
        },
        beforeDestroy : function() {
            this.$parent.showTabbar = true;
        }
    });

    var app = Vue.extend({
        data : function() {
            return {
                showTabbar : true
            };
        }
    });

    var router = new VueRouter();

    router.map({
        '/travellist' : {
            component : travellist
        },
        '/tourguide' : {
            component : tourguide
        },
        '/edittravel/:travelId' : {
            component : edittravel
        },
    });

    router.start(app, '#app');
});
