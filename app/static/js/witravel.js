$(function() {

    Vue.config.delimiters = [ '[[', ']]' ];
    var path = $('#path').val();

    var travellist = Vue.extend({
        template : $('#travellist_view').html(),
        data : function() {
            return {
                travellist : []
            };
        },
        methods : {
            say : function(aa) {
                alert(aa);
            }
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
                    type : 'next', // prev
                    lastId : ''
                }, function(data) {

                    for (var i = 0; i < 10; i++) {
                        me.travellist.push({
                            id : 'id-' + i,
                            title : 'hello-' + i,
                            description : '基于源数据将元素或模板块重复数次。指令的值必须使用特定语法 alias (in|of) expression，为当前遍历的元素提供别名：'
                        });
                    }

                    me.$nextTick(function() {
                        listScroll.refresh();
                        listScroll.loadFinish();
                    });

                });
            }

            // bind the iscroll

            // 上拉提示框状态，有hide, pulling, loading
            listScroll.pullUpStatus = 'hide';

            // 是否需要去加载更多
            listScroll.pullUpLoad = false;

            listScroll.on('scroll', function() {

                console.log("y is :" + this.y + ",max y is :" + this.maxScrollY);

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

                console.log("scroll end. y is :" + this.y + ",max y is :" + this.maxScrollY);

                // 不是正在加载，才进行加载
                if (this.pullUpStatus !== 'loading') {

                    if (this.pullUpLoad) {
                        this.pullUpStatus = 'loading';
                        $('.pull_up_tips').removeClass('pull_status_hide').removeClass('pull_status_pulling').addClass('pull_status_loading');
                        this.refresh();
                        this.scrollTo(0, this.maxScrollY, 1000, IScroll.utils.ease.elastic);

                        // do ajax
                        loadData();
                    } else {
                        this.pullUpStatus = 'hide';
                        $('.pull_up_tips').removeClass('pull_status_loading').removeClass('pull_status_pulling').addClass('pull_status_hide');
                    }

                    this.pullUpLoad = false;
                }
            });

            loadData();
        }
    });

    var tourguide = Vue.extend({
        template : '<p>This is tourguide!</p>'
    });

    var app = Vue.extend({});

    var router = new VueRouter();

    router.map({
        '/travellist' : {
            component : travellist
        },
        '/tourguide' : {
            component : tourguide
        },
    });

    router.start(app, '#app');
});
