import 'package:flutter/material.dart';

class TextFormFieldCustom extends StatelessWidget {
  const TextFormFieldCustom({
    super.key,
    required this.hintText,
    required this.icon,
    this.suffix,
  });

  final String hintText;
  final IconData? icon;
  final Widget? suffix;

  @override
  Widget build(BuildContext context) {
    return TextFormField(
      maxLines: 1,
      textInputAction: TextInputAction.next,
      keyboardType: TextInputType.emailAddress,
      decoration: InputDecoration(
        hintText: hintText,
        hintStyle: TextStyle(color: Colors.grey),
        filled: true,
        fillColor: Colors.grey.shade100,
        prefixIcon: Icon(icon),
        prefixIconColor: Colors.grey,
        suffixIcon: suffix,
        enabled: true,
        focusedBorder: OutlineInputBorder(
          borderRadius: BorderRadius.circular(16),
          borderSide: BorderSide(color: Colors.grey.shade300, width: 1),
        ),
        enabledBorder: OutlineInputBorder(
          borderRadius: BorderRadius.circular(16),
          borderSide: BorderSide(color: Colors.grey.shade100, width: 1),
        ),
      ),
      validator: (value) {
        if (value!.isEmpty) {
          return "$hintText is required";
        }
        return null;
      },
    );
  }
}
