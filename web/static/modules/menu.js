layui.define(["table", "form"], function (exports) {
  var $ = layui.$,
    form = layui.form,
    dropdown = layui.dropdown,
    tree = layui.tree,
    admin = layui.admin,
    treeTable = layui.treeTable;

  var prefix = "/system";
  var formTree = [];
  var icons = [];

  $.getJSON("/static/icons.json", function (data) {
    icons = data;
  });

  // 初始化菜单树
  function initMenuTree() {
    admin.req({
      url: prefix + "/menu/tree",
      done: function (res) {
        console.log(res);
        formTree = [
          {
            title: "根目录",
            is_parent: true,
            id: 0,
            children: res.data,
          },
        ];
      },
    });
  }

  initMenuTree();

  var treeTableIns = treeTable.render({
    elem: "#menu-tree-table",
    url: prefix + "/menu/tree",
    tree: {
      customName: {
        name: "title",
        pid: "parent_id",
        isParent: "is_parent",
      }
    },
    cols: [[
      { field: 'id', title: 'ID', width: 80, sort: true, fixed: 'left' },
      { field: 'title', title: '菜单名称' },
      { field: 'path', title: '菜单路径' },
      { field: 'sort', title: '显示顺序' },
      {
        field: 'type', width: 180, title: '菜单类型', templet: function (d) {
          if (d.type === 0) {
            return '<span class="layui-badge layui-bg-green">目录</span>';
          } else if (d.type === 1) {
            return '<span class="layui-badge layui-bg-blue">菜单</span>';
          } else {
            return '<span class="layui-badge layui-bg-gray">按钮</span>';
          }
        }
      },
      {
        fixed: "right", title: "操作", width: 280, align: "center", templet: function (d) {
          return `<a class="layui-btn layui-btn-xs" lay-event="edit"><i class="layui-icon layui-icon-edit"></i>编辑</a>
            <a class="layui-btn layui-btn-xs" lay-event="add"><i class="layui-icon layui-icon-add-1"></i>添加子菜单</a>
            <a class="layui-btn layui-btn-xs layui-btn-danger" lay-event="del"><i class="layui-icon layui-icon-delete"></i>删除</a>`;
        }
      }
    ]],
  });

  form.on('radio(type)', function (data) {
    if (data.value == 1) {
      $("#menu-path-item").show();
    }
  });

  // 父级菜单选择
  dropdown.render({
    elem: "#parent_id",
    content: '<div id="parent-menu-tree"></div>',
    style: "width: 460px; padding: 20px 10px;",
    ready: function (elemPanel, elem) {
      tree.render({
        elem: "#parent-menu-tree",
        data: formTree,
        onlyIconControl: true, // 是否仅允许节点左侧图标控制展开收缩
        click: function (obj) {
          // 给 form parent_id 赋值
          form.val("menu-form", {
            parent_id: obj.data.title,
          });
          $("#parent_id").data("id", obj.data.id);
          // 关闭 dropdown
          dropdown.close("parent_id");
        },
      });
    },
  });

  // 菜单图标选择
  var menuIconHtml = `<div id="menu-icons" class="layui-form">
    <input type="text" name="" placeholder="文本框" class="layui-input">
    <div class="icon-list" style="margin-top: 10px; overflow-y: auto; overflow-x: hidden; height: 220px; padding: 10px;"></div>
  </div>`;

  // 菜单图标下拉框
  dropdown.render({
    elem: "#menu-icon-dropdown",
    content: menuIconHtml,
    style: "width: 460px; height: 300px; padding: 20px 10px;",
    ready: function (elemPanel, elem) {
      var html = `<ul class="layui-row layui-col-space10">
          ${icons
          .map(
            (icon) => `
          <li class="layui-col-xs6 layui-col-sm4 layui-col-md2" style="cursor: pointer; margin-top: 10px;" data-icon="${icon.fontclass}">
            <i class="layui-icon ${icon.fontclass}" style="font-size: 20px;"></i>
          </li>
          `
          )
          .join("")}
        </ul>`;
      $(".icon-list").html(html);

      // 选择图标
      $(".icon-list li").click(function () {
        form.val("menu-form", {
          icon: "layui-icon " + $(this).data("icon"),
        });
        dropdown.close("menu-icon-dropdown");
      });

      // 输入框搜索
      $("#menu-icons input").on("input", function () {
        var val = $(this).val();
        $(".icon-list li").each(function () {
          if ($(this).data("icon").indexOf(val) > -1) {
            $(this).show();
          } else {
            $(this).hide();
          }
        });
      });
    },
  });

  form.on('submit(add-menu-submit)', function (data) {
    var field = data.field; // 获取表单字段值
    field.parent_id = $("#parent_id").data("id");
    field.sort = parseInt(field.sort);
    field.type = parseInt(field.type);
    var jsonData = JSON.stringify(field);
    admin.req({
      url: prefix + "/menu",
      method: "POST",
      contentType: "application/json",
      data: jsonData,
      done: function (res) {
        // 重新获取菜单树
        initMenuTree();
        treeTableIns.reload();
        // 清空表单
        $("#add-menu-reset").click();
        // 关闭弹出层
        layer.close("add");
        layer.msg("添加成功");
      },
    });
    return false;
  });

  // 按钮事件
  var active = {
    add: function () {
      layer.open({
        id: "add",
        type: 1,
        title: "添加菜单",
        content: $("#add"),
        area: ["600px", "600px"],
        btn: ["确定", "取消"],
        yes: function (index, layero) {
          var submit = layero.find("#add-menu-submit");
          submit.trigger("click");
        },
      });
    },
  };

  $(".layui-btn.layuiadmin-btn-admin").on("click", function () {
    var type = $(this).data("type");
    active[type] ? active[type].call(this) : "";
  });

  exports("menu", {});
});
