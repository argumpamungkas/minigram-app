import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:minigram/config/routes.dart';

void main() {
  runApp(MainApp());
}

class MainApp extends StatelessWidget {
  const MainApp({super.key});

  @override
  Widget build(BuildContext context) {
    return ScreenUtilInit(
      designSize: Size(375, 812),
      minTextAdapt: true,
      splitScreenMode: true,
      builder: (context, child) => MaterialApp(
        title: "Minigram",
        locale: Locale("id_ID"),
        theme: ThemeData(scaffoldBackgroundColor: Colors.grey.shade50),
        themeMode: ThemeMode.light,
        initialRoute: Routes.splashScreen,
        routes: Routes.routerApp,
      ),
    );
  }
}
