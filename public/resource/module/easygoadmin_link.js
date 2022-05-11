// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ 赋能开发者，助力企业发展 ]
// +----------------------------------------------------------------------
// | 版权所有 2019~2022 深圳EasyGoAdmin研发中心
// +----------------------------------------------------------------------
// | Licensed LGPL-3.0 EasyGoAdmin并不是自由软件，未经许可禁止去掉相关版权
// +----------------------------------------------------------------------
// | 官方网站: http://www.easygoadmin.vip
// +----------------------------------------------------------------------
// | Author: @半城风雨 团队荣誉出品 团队荣誉出品
// +----------------------------------------------------------------------
// | 版权和免责声明:
// | 本团队对该软件框架产品拥有知识产权（包括但不限于商标权、专利权、著作权、商业秘密等）
// | 均受到相关法律法规的保护，任何个人、组织和单位不得在未经本团队书面授权的情况下对所授权
// | 软件框架产品本身申请相关的知识产权，禁止用于任何违法、侵害他人合法权益等恶意的行为，禁
// | 止用于任何违反我国法律法规的一切项目研发，任何个人、组织和单位用于项目研发而产生的任何
// | 意外、疏忽、合约毁坏、诽谤、版权或知识产权侵犯及其造成的损失 (包括但不限于直接、间接、
// | 附带或衍生的损失等)，本团队不承担任何法律责任，本软件框架禁止任何单位和个人、组织用于
// | 任何违法、侵害他人合法利益等恶意的行为，如有发现违规、违法的犯罪行为，本团队将无条件配
// | 合公安机关调查取证同时保留一切以法律手段起诉的权利，本软件框架只能用于公司和个人内部的
// | 法律所允许的合法合规的软件产品研发，详细声明内容请阅读《框架免责声明》附件；
// +----------------------------------------------------------------------

/**
 * 友链管理
 * @author 半城风雨
 * @since 2021/7/26
 */
layui.use(['func'], function () {

    //声明变量
    var func = layui.func
        ,form = layui.form
        , $ = layui.$;

    if (A == 'index') {
        //【TABLE列数组】
        var cols = [
            {type: 'checkbox', fixed: 'left'}
            , {field: 'id', width: 80, title: 'ID', align: 'center', sort: true, fixed: 'left'}
            , {field: 'name', width: 250, title: '友链名称', align: 'center'}
            , {field: 'image', width: 100, title: '友链图片', align: 'center', templet: function (d) {
                    if (d.image != "") {
                        return '<a href="' + d.image + '" target="_blank"><img src="' + d.image + '" height="26" /></a>';
                    }
                }
            }
            , {field: 'type', width: 100, title: '类型', align: 'center', templet(d) {
                    var cls = "";
                    if (d.type == 1) {
                        // 友情链接
                        cls = "layui-btn-normal";
                    } else if (d.type == 2) {
                        // 合作伙伴
                        cls = "layui-btn-danger";
                    }
                    return '<span class="layui-btn ' + cls + ' layui-btn-xs">' + d.typeName + '</span>';
                }
            }
            , {field: 'url', width: 200, title: '友链地址', align: 'center', templet(d) {
                    return "<a href='" + d.url + "' target='_blank'>" + d.url + "</a>";
                }
            }
            , {field: 'platform', width: 100, title: '平台', align: 'center', templet(d) {
                    var cls = "";
                    if (d.platform == 1) {
                        // PC站
                        cls = "layui-btn-normal";
                    } else if (d.platform == 2) {
                        // WAP站
                        cls = "layui-btn-danger";
                    } else if (d.platform == 3) {
                        // 微信小程序
                        cls = "layui-btn-warm";
                    } else if (d.platform == 4) {
                        // APP应用
                        cls = "layui-btn-primary";
                    }
                    return '<span class="layui-btn ' + cls + ' layui-btn-xs">' + d.platformName + '</span>';
                }
            }
            , {field: 'form', width: 100, title: '友链形式', align: 'center', templet(d) {
                    var cls = "";
                    if (d.form == 1) {
                        // 文字链接
                        cls = "layui-btn-normal";
                    } else if (d.form == 2) {
                        // 图片链接
                        cls = "layui-btn-danger";
                    }
                    return '<span class="layui-btn ' + cls + ' layui-btn-xs">' + d.formName + '</span>';
                }
            }
            , {field: 'image', width: 100, title: '友链图片', align: 'center', templet: function (d) {
                    var imageStr = "";
                    if (d.imageUrl) {
                        imageStr = '<a href="' + d.imageUrl + '" target="_blank"><img src="' + d.imageUrl + '" height="26" /></a>';
                    }
                    return imageStr;
                }
            }
            , {field: 'status', width: 100, title: '状态', align: 'center', templet: function (d) {
                    return '<input type="checkbox" name="status" value="' + d.id + '" lay-skin="switch" lay-text="正常|禁用" lay-filter="status" ' + (d.status == 1 ? 'checked' : '') + '>';
                }
            }
            , {field: 'sort', width: 100, title: '显示顺序', align: 'center'}
            , {field: 'create_time', width: 180, title: '添加时间', align: 'center', templet:"<div>{{layui.util.toDateString(d.create_time*1000, 'yyyy-MM-dd HH:mm:ss')}}</div>"}
            , {field: 'update_time', width: 180, title: '更新时间', align: 'center', templet:"<div>{{layui.util.toDateString(d.update_time*1000, 'yyyy-MM-dd HH:mm:ss')}}</div>"}
            , {fixed: 'right', width: 150, title: '功能操作', align: 'center', toolbar: '#toolBar'}
        ];

        //【渲染TABLE】
        func.tableIns(cols, "tableList");

        //【设置弹框】
        func.setWin("友链");

        //【设置状态】
        func.formSwitch('status', null, function (data, res) {
            console.log("开关回调成功");
        });
    } else {
        //【监听友链类型】
        var link_form = $('#form').val();
        if (link_form == 1) {
            //文字
            $(".image").addClass("layui-hide");
        } else if (link_form == 2) {
            //图片
            $(".image").removeClass("layui-hide");
        }
        form.on('select(form)', function (data) {
            if (data.value == 1) {
                $(".image").addClass("layui-hide");
            } else if (data.value == 2) {
                $(".image").removeClass("layui-hide");
            }
        });
    }
});
